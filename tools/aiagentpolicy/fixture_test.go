package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func repoRoot(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("caller unavailable")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(file), "../.."))
}

func copyFixture(t *testing.T) string {
	t.Helper()
	root := repoRoot(t)
	dst := t.TempDir()
	for _, path := range []string{
		defaultManifestPath,
		"docs/20-domain/ai-agent-policy.md",
		".github/workflows/ai-agent-policy.yml",
	} {
		copyFile(t, filepath.Join(root, path), filepath.Join(dst, path))
	}
	m, err := loadManifest(filepath.Join(dst, defaultManifestPath))
	if err != nil {
		t.Fatal(err)
	}
	for _, reader := range m.RequiredGeneratedReaders {
		copyFile(t, filepath.Join(root, "docs/20-domain", reader), filepath.Join(dst, "docs/20-domain", reader))
	}
	return dst
}

func copyFile(t *testing.T, src, dst string) {
	t.Helper()
	body, err := os.ReadFile(src)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(dst, body, 0o644); err != nil {
		t.Fatal(err)
	}
}

func writeJSON(t *testing.T, path string, value manifest) {
	t.Helper()
	body, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, append(body, '\n'), 0o644); err != nil {
		t.Fatal(err)
	}
}

func mustContain(t *testing.T, body, want string) {
	t.Helper()
	if !strings.Contains(body, want) {
		t.Fatalf("missing %q in %s", want, body)
	}
}
