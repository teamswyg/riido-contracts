package main

import (
	"strings"
	"testing"
)

func TestBuildAPIProjectionReader(t *testing.T) {
	root, err := resolveRoot(".")
	if err != nil {
		t.Fatal(err)
	}
	m, summaries, doc, err := build(root, options{manifest: defaultManifest})
	if err != nil {
		t.Fatal(err)
	}
	if m.ID != "api-contract-projection" || len(summaries) != 2 {
		t.Fatalf("unexpected build result: id=%q summaries=%d", m.ID, len(summaries))
	}
	for _, phrase := range []string{
		"OpenAPI is not the SSOT",
		"Generated client delivery PRs are review handoffs",
		"riido.v2.aiAgent.tasks.assignedAgentProfiles",
		"generated client paths under `riido.v2.aiAgent.*`",
	} {
		if !strings.Contains(doc, phrase) {
			t.Fatalf("generated doc missing %q", phrase)
		}
	}
}

func TestAPIProjectionEvidenceCounts(t *testing.T) {
	summaries := []fixtureSummary{{OperationCount: 2, GeneratedPathCount: 1}}
	got := newEvidence(manifest{ID: "id", RequiredGeneratedPaths: []string{"x"}}, summaries)
	if got.Status != "verified" || got.OperationCount != 2 || got.GeneratedPathCount != 1 {
		t.Fatalf("evidence = %#v", got)
	}
}
