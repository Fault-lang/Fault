package listener

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	ospath "path/filepath"
	"strings"
)

// faultCacheSubdir is the subdirectory under the user's home directory used
// to cache remotely fetched specs.  Defined as a constant so it can be
// referenced in tests and future configuration paths without duplication.
const faultCacheSubdir = ".fault/cache"

// isURL reports whether rawPath is a valid HTTP or HTTPS URL.
// It uses url.Parse so that malformed strings (e.g. "http:///foo") fail early
// rather than being passed to http.Get with a confusing downstream error.
func isURL(rawPath string) bool {
	u, err := url.Parse(rawPath)
	if err != nil {
		return false
	}
	return (u.Scheme == "http" || u.Scheme == "https") && u.Host != ""
}

// faultCacheDir returns the path to the fault cache directory (~/.fault/cache).
// The directory is created if it does not already exist.
func faultCacheDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("cannot determine home directory: %w", err)
	}
	dir := ospath.Join(home, faultCacheSubdir)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", fmt.Errorf("cannot create cache directory %s: %w", dir, err)
	}
	return dir, nil
}

// urlCachePath returns the local cache file path for a given URL.
// The filename is the hex-encoded SHA-256 of the full URL so that it is
// collision-free regardless of path structure and avoids any character that is
// special to Fault or the host filesystem.
func urlCachePath(rawURL string) (string, error) {
	cacheDir, err := faultCacheDir()
	if err != nil {
		return "", err
	}
	_, err = url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("invalid URL %q: %w", rawURL, err)
	}
	// Use the full URL as the cache key so that two URLs that differ only in
	// path components always produce different filenames.
	cacheKey := fmt.Sprintf("%x", sha256.Sum256([]byte(rawURL)))
	return ospath.Join(cacheDir, cacheKey), nil
}

// fetchOrCacheURL returns the content of the spec at rawURL.  The result is
// cached in ~/.fault/cache/ so that subsequent runs (and offline use) do not
// require a network request.
//
// Set the environment variable FAULT_CACHE_REFRESH=1 to bypass the cache and
// re-fetch the spec from the network, replacing the cached copy.
func fetchOrCacheURL(rawURL string) ([]byte, error) {
	cachePath, err := urlCachePath(rawURL)
	if err != nil {
		return nil, err
	}

	// Return cached copy unless the caller requested a refresh.
	if os.Getenv("FAULT_CACHE_REFRESH") != "1" {
		if data, err := os.ReadFile(cachePath); err == nil {
			return data, nil
		}
	}

	// Fetch from the network.
	resp, err := http.Get(rawURL) // #nosec G107 -- URL is supplied by the Fault spec author
	if err != nil {
		return nil, fmt.Errorf("cannot fetch %q: %w", rawURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("cannot fetch %q: HTTP %d", rawURL, resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read response body from %q: %w", rawURL, err)
	}

	// Validate that the response looks like a Fault spec before caching.
	// A spec file must start with "spec " or "system " (after optional UTF-8 BOM).
	trimmed := strings.TrimPrefix(string(data), "\xef\xbb\xbf") // strip BOM if present
	trimmed = strings.TrimSpace(trimmed)
	if !strings.HasPrefix(trimmed, "spec ") && !strings.HasPrefix(trimmed, "system ") {
		return nil, fmt.Errorf("response from %q does not appear to be a Fault spec (unexpected content)", rawURL)
	}

	// Persist to cache using a restricted mode (owner read/write only).
	// A write failure (e.g. read-only filesystem, full disk) is silently ignored:
	// the data was successfully fetched and validated, so the import can proceed.
	_ = os.WriteFile(cachePath, data, 0o600)

	return data, nil
}
