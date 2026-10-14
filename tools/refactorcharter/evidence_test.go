package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestRunWritesEvidence(t *testing.T) {
	path := filepath.Join(t.TempDir(), "evidence.json")
	var out bytes.Buffer
	if err := run([]string{"-evidence-out", path}, &out); err != nil {
		t.Fatalf("run: %v", err)
	}
	body, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read evidence: %v", err)
	}
	var got evidence
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("decode evidence: %v", err)
	}
	if got.SchemaVersion != evidenceSchemaVersion {
		t.Fatalf("unexpected evidence: %+v", got)
	}
	if got.FilesScanned == 0 || got.TargetMaxLines == 0 {
		t.Fatalf("missing scan evidence: %+v", got)
	}
	if got.Status != evidenceStatus(charter{Mode: got.Mode}, scanReport{Findings: got.Findings}) {
		t.Fatalf("status does not match findings: %+v", got)
	}
}
