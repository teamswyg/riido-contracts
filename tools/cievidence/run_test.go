package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestRunWritesVerifiedEvidence(t *testing.T) {
	dir := t.TempDir()
	workflow := filepath.Join(dir, ".github", "workflows", "ci.yml")
	body := `run: test "$(go list -m all | wc -l | tr -d ' ')" = "1"
run: go run ./tools/enumgen verify
run: go run ./tools/fsmgen verify
run: go test ./...
`
	if err := os.MkdirAll(filepath.Dir(workflow), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(workflow, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	out := filepath.Join(dir, "out", "ci.json")
	err := run([]string{"-workflow", workflow, "-id", "ci", "-evidence-out", out})
	if err != nil {
		t.Fatalf("run: %v", err)
	}
	var got evidence
	data, err := os.ReadFile(out)
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}
	if got.Status != "verified" || len(got.Commands) != 4 {
		t.Fatalf("unexpected evidence: %+v", got)
	}
}
