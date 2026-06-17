package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDefaultCharterRunsInAdvisoryMode(t *testing.T) {
	var out bytes.Buffer
	if err := run(nil, &out); err != nil {
		t.Fatalf("run default charter: %v", err)
	}
	if got := out.String(); !strings.Contains(got, "mode=advisory") {
		t.Fatalf("unexpected output: %q", got)
	}
}

func TestEnforcedModeFailsOnLongFile(t *testing.T) {
	root := t.TempDir()
	mustWrite(t, root, "src/long.go", strings.Repeat("x\n", 4))
	manifest := `{
		"schema_version":"riido-refactoring-charter.v1",
		"id":"test-charter",
		"riido_task":"RIID-TEST",
		"mode":"enforced",
		"line_budget":{"target_max_lines":3,"recommended_min_lines":1,"recommended_max_lines":2},
		"semantic_units":["concept"],
		"required_artifacts":["verification"],
		"scan":{"roots":["src"],"include_extensions":[".go"],"generated_path_fragments":[],"generated_markers":[]}
	}`
	manifestPath := filepath.Join(root, "charter.json")
	if err := os.WriteFile(manifestPath, []byte(manifest), 0o644); err != nil {
		t.Fatal(err)
	}
	err := run([]string{"-manifest", manifestPath, "-root", root}, ioDiscard{})
	if err == nil || !strings.Contains(err.Error(), "files exceed") {
		t.Fatalf("expected enforced failure, got %v", err)
	}
}

func mustWrite(t *testing.T, root, rel, body string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}

type ioDiscard struct{}

func (ioDiscard) Write(p []byte) (int, error) { return len(p), nil }
