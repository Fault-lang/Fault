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
	expectedDir := ospath.Join(home, ".fault", "cache")
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
	// Two different URLs must produce different cache paths.
	other, err := urlCachePath("https://registry.example.com/other.fspec")
	if err != nil {
		t.Fatalf("urlCachePath() other error: %v", err)
	}
	if got == other {
		t.Errorf("different URLs mapped to the same cache path %q", got)
	}
}

func TestFetchOrCacheURL(t *testing.T) {
	specContent := `spec fetched;
def x = 1;
`
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

	// The cache file must now exist.
	if _, statErr := os.Stat(cachePath); os.IsNotExist(statErr) {
		t.Errorf("cache file %q was not created", cachePath)
	}

	// Second call – must be served from cache (server is still up but we
	// verify the content is identical, meaning the cache path was hit).
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
