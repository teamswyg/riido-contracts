package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestVerifyWritesEvidence(t *testing.T) {
	root := repoRoot(t)
	out := filepath.Join(t.TempDir(), "evidence.json")
	err := run([]string{"-root", root, "-check-doc", "-evidence-out", out})
	if err != nil {
		t.Fatal(err)
	}
	body, err := os.ReadFile(out)
	if err != nil {
		t.Fatal(err)
	}
	mustContain(t, string(body), `"status": "verified"`)
	mustContain(t, string(body), `"policy_assertion_count": 13`)
}

func TestVerifierRejectsSectionHeadingLeak(t *testing.T) {
	root := copyFixture(t)
	path := filepath.Join(root, defaultManifestPath)
	m, err := loadManifest(path)
	if err != nil {
		t.Fatal(err)
	}
	m.Sections[0].Body = append(m.Sections[0].Body, "## Boundary")
	writeJSON(t, path, m)
	if err := run([]string{"-root", root}); err == nil {
		t.Fatal("expected heading leak to fail")
	}
}
