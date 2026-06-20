package main

import "testing"

func TestBuildAIAgentAPISurfaceEvidence(t *testing.T) {
	root, err := resolveRoot(".")
	if err != nil {
		t.Fatal(err)
	}
	model, doc, err := build(root, options{root: ".", manifest: defaultManifestPath})
	if err != nil {
		t.Fatal(err)
	}
	if len(model.V2Only) != model.Manifest.ExpectedV2OnlyOperations {
		t.Fatalf("v2 only = %d", len(model.V2Only))
	}
	if len(doc) == 0 {
		t.Fatal("rendered doc is empty")
	}
}
