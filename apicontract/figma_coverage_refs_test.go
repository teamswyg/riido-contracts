package apicontract

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func assertCoverageLocalRefExists(t *testing.T, ref string) {
	t.Helper()
	path := ref
	if before, _, ok := strings.Cut(ref, "#"); ok {
		path = before
	}
	if strings.TrimSpace(path) == "" {
		t.Fatalf("empty local ref in %q", ref)
	}
	if _, err := os.Stat(filepath.FromSlash("../" + path)); err != nil {
		t.Fatalf("local ref %q does not exist: %v", ref, err)
	}
}
