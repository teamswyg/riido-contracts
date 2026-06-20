package main

import (
	"strings"
	"testing"
)

func TestBuildProviderCapabilityReader(t *testing.T) {
	root, err := resolveRoot(".")
	if err != nil {
		t.Fatal(err)
	}
	model, doc, err := build(root, options{manifest: defaultManifest})
	if err != nil {
		t.Fatal(err)
	}
	if len(model.Protocols) != 6 || model.CriticalArgSetCount != 5 {
		t.Fatalf("unexpected model: %#v", model)
	}
	for _, phrase := range []string{"claude-stream-json", "CapabilityFingerprint", "DetectedVersion"} {
		if !strings.Contains(doc, phrase) {
			t.Fatalf("generated doc missing %q", phrase)
		}
	}
}

func TestEvidenceCounts(t *testing.T) {
	model := model{Manifest: manifest{ID: "id", Invariants: []string{"x"}}, Protocols: []protocolRow{{}}, CriticalArgSetCount: 1}
	got := newEvidence(model)
	if got.Status != "verified" || got.ProtocolCount != 1 || got.InvariantCount != 1 {
		t.Fatalf("evidence = %#v", got)
	}
}
