package main

import (
	"strings"
	"testing"
)

func TestBuildIREventLogReader(t *testing.T) {
	root, err := resolveRoot(".")
	if err != nil {
		t.Fatal(err)
	}
	model, doc, err := build(root, options{manifest: defaultManifest})
	if err != nil {
		t.Fatal(err)
	}
	if model.EventCount != 62 || model.TransitionCount != 23 {
		t.Fatalf("unexpected model counts: %#v", model)
	}
	for _, phrase := range []string{"TaskQueued", "Reducer Surface", "Native Config Classification"} {
		if !strings.Contains(doc, phrase) {
			t.Fatalf("generated doc missing %q", phrase)
		}
	}
}

func TestEvidenceCounts(t *testing.T) {
	model := model{Manifest: manifest{ID: "id"}, EventCount: 2, TransitionCount: 1}
	got := newEvidence(model)
	if got.Status != "verified" || got.EventCount != 2 || got.TransitionCount != 1 {
		t.Fatalf("evidence = %#v", got)
	}
}
