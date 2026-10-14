package main

import "testing"

func TestBuildHostIntegrationEvidence(t *testing.T) {
	root, err := resolveRoot(".")
	if err != nil {
		t.Fatal(err)
	}
	model, doc, err := build(root, options{root: ".", manifest: defaultManifestPath})
	if err != nil {
		t.Fatal(err)
	}
	if !model.StoreManagedExclusive {
		t.Fatal("store-managed classification is not exclusive")
	}
	if len(doc) == 0 {
		t.Fatal("rendered doc is empty")
	}
}
