package listener

import (
	"net/http"
	"net/http/httptest"
	"os"
	ospath "path/filepath"
	"testing"
)

func TestIsURL(t *testing.T) {
	cases := []struct {
		path string
		want bool
	}{
		{"http://example.com/spec.fspec", true},
		{"https://example.com/spec.fspec", true},
		{"./local/spec.fspec", false},
		{"../relative/spec.fspec", false},
		{"/absolute/spec.fspec", false},
		{"ftp://example.com/spec.fspec", false},
		// Malformed URLs that start with "http://" but have no host must be rejected.
		{"http:///foo", false},
		{"https:///bar", false},
	}
	for _, tc := range cases {
		if got := isURL(tc.path); got != tc.want {
			t.Errorf("isURL(%q) = %v, want %v", tc.path, got, tc.want)
		}
	}
}

func TestURLCachePath(t *testing.T) {
	rawURL := "https://registry.example.com/law-specs/flsa/v2025-01-01.fspec"
	got, err := urlCachePath(rawURL)
	if err != nil {
		t.Fatalf("urlCachePath() error: %v", err)
	}
	if got == "" {
		t.Fatal("urlCachePath() returned empty string")
	}
	// The cache file must live inside ~/.fault/cache/.
	home, _ := os.UserHomeDir()
	expectedDir := ospath.Join(home, faultCacheSubdir)
	if dir := ospath.Dir(got); dir != expectedDir {
		t.Errorf("cache path directory = %q, want %q", dir, expectedDir)
	}
	// Two calls for the same URL must return the same path.
	got2, err := urlCachePath(rawURL)
	if err != nil {
		t.Fatalf("urlCachePath() second call error: %v", err)
	}
	if got != got2 {
		t.Errorf("urlCachePath() not deterministic: %q != %q", got, got2)
	}
	// Two different URLs that share a common prefix must produce different cache paths.
	other, err := urlCachePath("https://registry.example.com/other.fspec")
	if err != nil {
		t.Fatalf("urlCachePath() other error: %v", err)
	}
	if got == other {
		t.Errorf("different URLs mapped to the same cache path %q", got)
	}
	// Paths that differ only by _ vs / must also differ (collision-free hashing).
	collision1, _ := urlCachePath("https://host/a_b")
	collision2, _ := urlCachePath("https://host/a/b")
	if collision1 == collision2 {
		t.Errorf("path collision: host/a_b and host/a/b mapped to the same key")
	}
}

func TestFetchOrCacheURL(t *testing.T) {
	specContent := "spec fetched;\ndef x = 1;\n"
	// Start a local HTTP server that serves the spec content.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(specContent))
	}))
	defer ts.Close()

	rawURL := ts.URL + "/law-specs/test.fspec"

	// Clean up any cached copy from a prior run.
	cachePath, err := urlCachePath(rawURL)
	if err != nil {
		t.Fatalf("urlCachePath() error: %v", err)
	}
	_ = os.Remove(cachePath)

	// First call – fetches from the test server.
	data, err := fetchOrCacheURL(rawURL)
	if err != nil {
		t.Fatalf("fetchOrCacheURL() first call error: %v", err)
	}
	if string(data) != specContent {
		t.Errorf("fetchOrCacheURL() content = %q, want %q", string(data), specContent)
	}

	// The cache file must now exist with restricted permissions (0600).
	fi, statErr := os.Stat(cachePath)
	if os.IsNotExist(statErr) {
		t.Fatalf("cache file %q was not created", cachePath)
	}
	if fi.Mode().Perm() != 0o600 {
		t.Errorf("cache file permissions = %o, want 0600", fi.Mode().Perm())
	}

	// Second call – must be served from cache.
	data2, err := fetchOrCacheURL(rawURL)
	if err != nil {
		t.Fatalf("fetchOrCacheURL() second call error: %v", err)
	}
	if string(data2) != specContent {
		t.Errorf("fetchOrCacheURL() cached content = %q, want %q", string(data2), specContent)
	}

	// Clean up.
	_ = os.Remove(cachePath)
}

func TestFetchOrCacheURL_Refresh(t *testing.T) {
	calls := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("spec refreshed;\n"))
	}))
	defer ts.Close()

	rawURL := ts.URL + "/refresh.fspec"
	cachePath, _ := urlCachePath(rawURL)
	_ = os.Remove(cachePath)

	// First fetch populates the cache.
	if _, err := fetchOrCacheURL(rawURL); err != nil {
		t.Fatalf("first fetch: %v", err)
	}
	if calls != 1 {
		t.Fatalf("expected 1 network call after first fetch, got %d", calls)
	}

	// Second call without refresh flag must use the cache.
	if _, err := fetchOrCacheURL(rawURL); err != nil {
		t.Fatalf("second fetch: %v", err)
	}
	if calls != 1 {
		t.Errorf("expected still 1 network call after cached fetch, got %d", calls)
	}

	// Setting FAULT_CACHE_REFRESH=1 must bypass the cache.
	t.Setenv("FAULT_CACHE_REFRESH", "1")
	if _, err := fetchOrCacheURL(rawURL); err != nil {
		t.Fatalf("refresh fetch: %v", err)
	}
	if calls != 2 {
		t.Errorf("expected 2 network calls after forced refresh, got %d", calls)
	}

	_ = os.Remove(cachePath)
}

func TestFetchOrCacheURL_HTTPError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer ts.Close()

	rawURL := ts.URL + "/missing.fspec"

	// Clean up any cached copy.
	cachePath, _ := urlCachePath(rawURL)
	_ = os.Remove(cachePath)

	_, err := fetchOrCacheURL(rawURL)
	if err == nil {
		t.Fatal("fetchOrCacheURL() expected error for HTTP 404, got nil")
	}
}

func TestFetchOrCacheURL_NotASpec(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("<html><body>Not a spec</body></html>"))
	}))
	defer ts.Close()

	rawURL := ts.URL + "/html-page.fspec"
	cachePath, _ := urlCachePath(rawURL)
	_ = os.Remove(cachePath)

	_, err := fetchOrCacheURL(rawURL)
	if err == nil {
		t.Fatal("fetchOrCacheURL() expected error for HTML response, got nil")
	}

	// The bad response must not have been cached.
	if _, statErr := os.Stat(cachePath); !os.IsNotExist(statErr) {
		t.Errorf("cache file %q should not exist after failed content validation", cachePath)
		_ = os.Remove(cachePath)
	}
}

func TestImportURLSyntax(t *testing.T) {
	// In testing mode the listener does not perform actual I/O, so we just
	// verify that an import statement containing an HTTPS URL is parsed into
	// the AST with the correct path and identifier.
	test := `system test1;
			 import flsa "https://registry.example.com/law-specs/flsa/v2025-01-01.fspec";
			`
	flags := make(map[string]bool)
	flags["specType"] = false
	_, spec := prepTest(test, flags)
	if spec == nil {
		t.Fatalf("prepTest() returned nil")
	}
	if len(spec.Statements) != 2 {
		t.Fatalf("spec.Statements does not contain 2 statements. got=%d", len(spec.Statements))
	}

	if spec.Statements[1].TokenLiteral() != "IMPORT_DECL" {
		t.Fatalf("spec.Statement[1] is not an import statement. got=%T", spec.Statements[1])
	}
}
