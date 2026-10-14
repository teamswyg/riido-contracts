package main

import "testing"

func TestBuildDevicePrincipalEvidence(t *testing.T) {
	root, err := resolveRoot(".")
	if err != nil {
		t.Fatal(err)
	}
	model, doc, err := build(root, options{root: ".", manifest: defaultManifestPath})
	if err != nil {
		t.Fatal(err)
	}
	if model.DependencyPhraseCount != model.Manifest.ExpectedDependencyPhraseCount {
		t.Fatalf("dependency phrase count = %d", model.DependencyPhraseCount)
	}
	if len(doc) == 0 {
		t.Fatal("rendered doc is empty")
	}
}
