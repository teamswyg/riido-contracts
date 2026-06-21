package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestRunRejectsManualReaderAndWritesEvidence(t *testing.T) {
	root := fixtureRepo(t)
	out := filepath.Join(root, "out", "coverage.json")
	err := run([]string{"-root", root, "-write-doc", "-check-doc", "-evidence-out", out})
	if err == nil {
		t.Fatal("expected manual reader failure")
	}
	body, err := os.ReadFile(out)
	if err != nil {
		t.Fatalf("read evidence: %v", err)
	}
	var got evidence
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("decode evidence: %v", err)
	}
	if got.Status != "failed" || got.ManualCount != 1 || got.GeneratedCount != 1 {
		t.Fatalf("unexpected evidence: %+v", got)
	}
	if got.ManifestInventory != 1 || len(got.ManifestInventoryByGroup) != 1 {
		t.Fatalf("missing manifest inventory breakdown: %+v", got)
	}
	if len(got.ManifestInventorySamples) != 1 || len(got.ManifestInventorySamples[0].Paths) != 1 {
		t.Fatalf("missing manifest inventory samples: %+v", got)
	}
}

func TestCheckDocRejectsStaleReader(t *testing.T) {
	root := fixtureRepo(t)
	if err := writeFile(filepath.Join(root, "docs", "executable-knowledge.md"), []byte("stale")); err != nil {
		t.Fatal(err)
	}
	err := run([]string{"-root", root, "-check-doc"})
	if err == nil {
		t.Fatal("expected stale doc failure")
	}
}
