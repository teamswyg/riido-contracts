package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestVerifyWritesEvidence(t *testing.T) {
	path := filepath.Join(t.TempDir(), "evidence.json")
	var out bytes.Buffer
	if err := run([]string{"verify", "-evidence-out", path}, &out); err != nil {
		t.Fatalf("verify: %v", err)
	}
	body, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read evidence: %v", err)
	}
	var got evidence
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("decode evidence: %v", err)
	}
	if got.SchemaVersion != evidenceSchemaVersion || got.Status != "verified" {
		t.Fatalf("unexpected evidence: %+v", got)
	}
	if got.EntriesVerified == 0 || got.GeneratedAnnotationsChecked == 0 {
		t.Fatalf("missing evidence counts: %+v", got)
	}
}
