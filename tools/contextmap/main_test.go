package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRenderManifestMarksGeneratedDoc(t *testing.T) {
	body := renderManifest(minimalManifest())
	if !strings.Contains(body, generatedMarker) {
		t.Fatalf("missing generated marker: %s", body)
	}
	if !strings.Contains(body, "## Owned Contexts") || !strings.Contains(body, "## Non-Owned Contexts") {
		t.Fatalf("missing context sections: %s", body)
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
	if got.Status != "verified" || got.OwnedContextCount == 0 || got.SSOTLinkCount == 0 {
		t.Fatalf("unexpected evidence: %+v", got)
	}
}
