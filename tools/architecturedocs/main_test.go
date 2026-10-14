package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRenderDocsAreGenerated(t *testing.T) {
	m := minimalManifest()
	if !strings.Contains(renderModuleDoc(m), generatedMarker) {
		t.Fatal("module doc missing generated marker")
	}
	if !strings.Contains(renderIntegrationDoc(m), "go test ./...") {
		t.Fatal("integration doc missing command")
	}
}

func TestVerifyWritesEvidence(t *testing.T) {
	path := filepath.Join(t.TempDir(), "evidence.json")
	if err := run([]string{"verify", "-evidence-out", path}, os.Stdout); err != nil {
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
	if got.Status != "verified" || got.PackageCount == 0 || got.PublicGateCount == 0 {
		t.Fatalf("unexpected evidence: %+v", got)
	}
}
