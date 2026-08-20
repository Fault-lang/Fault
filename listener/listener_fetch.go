package listener

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	ospath "path/filepath"
	"strings"
)

// isURL reports whether rawPath is an HTTP or HTTPS URL.
func isURL(rawPath string) bool {
	return strings.HasPrefix(rawPath, "http://") || strings.HasPrefix(rawPath, "https://")
}

// faultCacheDir returns the path to the fault cache directory (~/.fault/cache).
// The directory is created if it does not already exist.
func faultCacheDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("cannot determine home directory: %w", err)
	}
	dir := ospath.Join(home, ".fault", "cache")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", fmt.Errorf("cannot create cache directory %s: %w", dir, err)
	}
	return dir, nil
}

// urlCachePath returns the local cache file path for a given URL.
// The path component of the URL is used as the filename, with directory
// separators replaced by underscores to produce a flat file layout.
func urlCachePath(rawURL string) (string, error) {
	cacheDir, err := faultCacheDir()
	if err != nil {
		return "", err
	}
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("invalid URL %q: %w", rawURL, err)
	}
	// Build a filename that is unique per (host, path).
	segment := u.Host + u.Path
	segment = strings.ReplaceAll(segment, "/", "_")
	segment = strings.ReplaceAll(segment, ":", "_")
	return ospath.Join(cacheDir, segment), nil
}

// fetchOrCacheURL returns the content of the spec at rawURL.  The result is
// cached in ~/.fault/cache/ so that subsequent runs (and offline use) do not
// require a network request.
func fetchOrCacheURL(rawURL string) ([]byte, error) {
	cachePath, err := urlCachePath(rawURL)
	if err != nil {
		return nil, err
	}

	// Return cached copy if it exists.
	if data, err := os.ReadFile(cachePath); err == nil {
		return data, nil
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

	// Persist to cache.
	if writeErr := os.WriteFile(cachePath, data, 0o644); writeErr != nil {
		// Non-fatal: we still have the data in memory.
		fmt.Fprintf(os.Stderr, "warning: could not cache %q at %s: %v\n", rawURL, cachePath, writeErr)
	}

	return data, nil
}
