package main

import "testing"

func TestBuildAIAgentConfigurationEvidence(t *testing.T) {
	root, err := resolveRoot(".")
	if err != nil {
		t.Fatal(err)
	}
	model, doc, err := build(root, options{root: ".", manifest: defaultManifestPath})
	if err != nil {
		t.Fatal(err)
	}
	if model.ScenarioCount != model.Manifest.ExpectedScenarioCount {
		t.Fatalf("scenario count = %d", model.ScenarioCount)
	}
	if len(doc) == 0 {
		t.Fatal("rendered doc is empty")
	}
}
