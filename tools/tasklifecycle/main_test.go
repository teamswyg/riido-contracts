package main

import (
	"strings"
	"testing"
)

func TestBuildTaskLifecycleReader(t *testing.T) {
	root, err := resolveRoot(".")
	if err != nil {
		t.Fatal(err)
	}
	model, doc, err := build(root, options{manifest: defaultManifest})
	if err != nil {
		t.Fatal(err)
	}
	if model.TransitionCount != 43 || len(model.States) != 15 {
		t.Fatalf("unexpected model counts: %#v", model)
	}
	for _, phrase := range []string{"Running", "RunReportedDone -> Validating", "Terminal states have no outgoing transitions"} {
		if !strings.Contains(doc, phrase) {
			t.Fatalf("generated doc missing %q", phrase)
		}
	}
}

func TestEvidenceCounts(t *testing.T) {
	model := model{Manifest: manifest{ID: "id", Invariants: []string{"x"}}, FSMSchema: 1, States: []stateRow{{}}, TransitionCount: 2}
	got := newEvidence(model)
	if got.Status != "verified" || got.States != 1 || got.Transitions != 2 || got.Invariants != 1 {
		t.Fatalf("evidence = %#v", got)
	}
}
