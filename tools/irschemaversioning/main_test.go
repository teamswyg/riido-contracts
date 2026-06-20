package main

import (
	"strings"
	"testing"
)

func TestBuildIRSchemaVersioningReader(t *testing.T) {
	root, err := resolveRoot(".")
	if err != nil {
		t.Fatal(err)
	}
	model, doc, err := build(root, options{manifest: defaultManifest})
	if err != nil {
		t.Fatal(err)
	}
	if model.CanonicalEventFields != 23 || model.EventScopeCount != 4 {
		t.Fatalf("unexpected model counts: %#v", model)
	}
	for _, phrase := range []string{"Scope Rules", "ValidateEnvelope", "Fake placeholder"} {
		if !strings.Contains(doc, phrase) {
			t.Fatalf("generated doc missing %q", phrase)
		}
	}
}

func TestEvidenceCounts(t *testing.T) {
	model := model{Manifest: manifest{ID: "id"}, CanonicalEventFields: 2, EventScopeCount: 1}
	got := newEvidence(model)
	if got.Status != "verified" || got.CanonicalEventFields != 2 || got.EventScopeCount != 1 {
		t.Fatalf("evidence = %#v", got)
	}
}
