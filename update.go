package main

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	ospath "path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	githubRepo    = "Fault-lang/Fault"
	updateCheckTTL = 24 * time.Hour
)

type updateCache struct {
	CheckedAt     time.Time `json:"checked_at"`
	LatestVersion string    `json:"latest_version"`
}

type githubRelease struct {
	TagName string        `json:"tag_name"`
	Assets  []githubAsset `json:"assets"`
}

type githubAsset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

// startupUpdateCheck prints a warning if a newer version of fault is available.
// It caches the result of GitHub API calls so the network is only hit once per day.
func startupUpdateCheck() {
	latest, hasUpdate, err := checkForUpdateCached()
	if err != nil || !hasUpdate {
		return
	}
	fmt.Fprintf(os.Stderr, "A new version of fault is available: %s (you have %s). Run `fault update` to upgrade.\n\n", latest, version)
}

func checkForUpdateCached() (latestVersion string, hasUpdate bool, err error) {
	home, herr := os.UserHomeDir()
	if herr != nil {
		// No home dir — just check directly
		latestVersion, err = fetchLatestVersion()
		if err == nil {
			hasUpdate = isNewerVersion(latestVersion, version)
		}
		return
	}

	cachePath := ospath.Join(home, ".fault_update_cache")

	// Try reading a fresh cache
	data, rerr := os.ReadFile(cachePath)
	if rerr == nil {
		var cache updateCache
		if jerr := json.Unmarshal(data, &cache); jerr == nil && time.Since(cache.CheckedAt) < updateCheckTTL {
			latestVersion = cache.LatestVersion
			hasUpdate = isNewerVersion(latestVersion, version)
			return
		}
	}

	// Cache is stale or missing — fetch from GitHub
	latestVersion, err = fetchLatestVersion()
	if err != nil {
		return
	}

	// Write updated cache (best-effort)
	if b, merr := json.Marshal(updateCache{CheckedAt: time.Now(), LatestVersion: latestVersion}); merr == nil {
		os.WriteFile(cachePath, b, 0644)
	}

	hasUpdate = isNewerVersion(latestVersion, version)
	return
}

func fetchLatestVersion() (string, error) {
	release, err := fetchRelease(3 * time.Second)
	if err != nil {
		return "", err
	}
	return release.TagName, nil
}

func fetchRelease(timeout time.Duration) (*githubRelease, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	url := "https://api.github.com/repos/" + githubRepo + "/releases/latest"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned %s", resp.Status)
	}

	var release githubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, err
	}
	return &release, nil
}

// runUpdate downloads and installs the latest release, replacing the current binary.
func runUpdate() error {
	fmt.Println("Checking for updates...")

	release, err := fetchRelease(30 * time.Second)
	if err != nil {
		return fmt.Errorf("could not check for updates: %w", err)
	}

	if !isNewerVersion(release.TagName, version) {
		fmt.Printf("Already up to date (%s).\n", version)
		return nil
	}

	fmt.Printf("Updating to %s (you have %s)...\n", release.TagName, version)

	assetName := platformAssetName()
	var downloadURL string
	for _, a := range release.Assets {
		if a.Name == assetName {
			downloadURL = a.BrowserDownloadURL
			break
		}
	}
	if downloadURL == "" {
		return fmt.Errorf("no release asset found for %s/%s (looked for %s)", runtime.GOOS, runtime.GOARCH, assetName)
	}

	fmt.Printf("Downloading %s...\n", assetName)

	tmp, err := os.CreateTemp("", "fault-update-*")
	if err != nil {
		return fmt.Errorf("could not create temp file: %w", err)
	}
	tmpPath := tmp.Name()
	defer os.Remove(tmpPath)

	if err := downloadTo(tmp, downloadURL); err != nil {
		tmp.Close()
		return fmt.Errorf("download failed: %w", err)
	}
	tmp.Close()

	binaryName := "fault"
	if runtime.GOOS == "windows" {
		binaryName = "fault.exe"
	}

	extractedPath, err := extractBinary(tmpPath, binaryName)
	if err != nil {
		return fmt.Errorf("extraction failed: %w", err)
	}
	defer os.Remove(extractedPath)

	execPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("could not determine current executable path: %w", err)
	}
	execPath, err = ospath.EvalSymlinks(execPath)
	if err != nil {
		return fmt.Errorf("could not resolve executable path: %w", err)
	}

	if err := replaceBinary(extractedPath, execPath); err != nil {
		return fmt.Errorf("could not install update (try running with sudo): %w", err)
	}

	fmt.Printf("Successfully updated to %s.\n", release.TagName)
	return nil
}

func platformAssetName() string {
	goos := strings.ToUpper(runtime.GOOS[:1]) + runtime.GOOS[1:] // Darwin, Linux, Windows

	var arch string
	switch runtime.GOARCH {
	case "amd64":
		arch = "x86_64"
	case "386":
		arch = "i386"
	default:
		arch = runtime.GOARCH
	}

	ext := ".tar.gz"
	if runtime.GOOS == "windows" {
		ext = ".zip"
	}

	return fmt.Sprintf("fault_%s_%s%s", goos, arch, ext)
}

func downloadTo(dst *os.File, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned %s", resp.Status)
	}

	_, err = io.Copy(dst, resp.Body)
	return err
}

// extractBinary extracts the named binary from a tar.gz or zip archive and
// writes it to a new temp file, returning its path.
func extractBinary(archivePath, binaryName string) (string, error) {
	tmp, err := os.CreateTemp("", "fault-binary-*")
	if err != nil {
		return "", err
	}
	outPath := tmp.Name()

	var extractErr error
	if runtime.GOOS == "windows" {
		extractErr = extractFromZip(archivePath, binaryName, tmp)
	} else {
		extractErr = extractFromTarGz(archivePath, binaryName, tmp)
	}
	tmp.Close()

	if extractErr != nil {
		os.Remove(outPath)
		return "", extractErr
	}
	return outPath, nil
}

func extractFromTarGz(archivePath, binaryName string, dst *os.File) error {
	f, err := os.Open(archivePath)
	if err != nil {
		return err
	}
	defer f.Close()

	gz, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer gz.Close()

	tr := tar.NewReader(gz)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if ospath.Base(hdr.Name) == binaryName {
			if _, err := io.Copy(dst, tr); err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("%s not found in archive", binaryName)
}

func extractFromZip(archivePath, binaryName string, dst *os.File) error {
	r, err := zip.OpenReader(archivePath)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		if ospath.Base(f.Name) == binaryName {
			rc, err := f.Open()
			if err != nil {
				return err
			}
			_, err = io.Copy(dst, rc)
			rc.Close()
			return err
		}
	}
	return fmt.Errorf("%s not found in archive", binaryName)
}

// replaceBinary atomically replaces the target binary with the new one.
// On Windows it renames the old binary out of the way first since you
// can't overwrite a running executable.
func replaceBinary(newPath, targetPath string) error {
	if err := os.Chmod(newPath, 0755); err != nil {
		return err
	}

	if runtime.GOOS == "windows" {
		oldPath := targetPath + ".old"
		os.Remove(oldPath) // clean up any leftover
		if err := os.Rename(targetPath, oldPath); err != nil {
			return err
		}
	}

	return os.Rename(newPath, targetPath)
}

// isNewerVersion returns true if latest is strictly greater than current.
// Both versions are expected in vX.Y.Z format.
func isNewerVersion(latest, current string) bool {
	l := parseVersion(latest)
	c := parseVersion(current)
	for i := range l {
		if i >= len(c) {
			return true
		}
		if l[i] > c[i] {
			return true
		}
		if l[i] < c[i] {
			return false
		}
	}
	return false
}

func parseVersion(v string) []int {
	v = strings.TrimPrefix(v, "v")
	parts := strings.Split(v, ".")
	result := make([]int, len(parts))
	for i, p := range parts {
		result[i], _ = strconv.Atoi(p)
	}
	return result
}
