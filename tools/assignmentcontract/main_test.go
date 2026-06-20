package main

import (
	"strings"
	"testing"
)

func TestBuildAssignmentContractReader(t *testing.T) {
	root, err := resolveRoot(".")
	if err != nil {
		t.Fatal(err)
	}
	m, c, doc, err := build(root, options{manifest: defaultManifest})
	if err != nil {
		t.Fatal(err)
	}
	if m.ID != "assignment-polling-contract" {
		t.Fatalf("manifest id = %q", m.ID)
	}
	if len(c.AssignmentStates) != 8 || len(c.PollActions) != 4 {
		t.Fatalf("unexpected contract counts: states=%d poll_actions=%d", len(c.AssignmentStates), len(c.PollActions))
	}
	for _, phrase := range []string{
		generatedNotice,
		"Assignment.allow_experimental_runtime",
		"assignment heartbeat every 5 seconds",
		"not been refreshed for 20 seconds",
	} {
		if !strings.Contains(doc, phrase) {
			t.Fatalf("generated doc missing %q", phrase)
		}
	}
}

func TestAssignmentContractEvidence(t *testing.T) {
	c := contract{SchemaVersion: "contract", ServiceSchemaVersion: "service"}
	c.AssignmentStates = []state{{Value: "queued"}}
	c.PollActions = []namedValue{{Value: "none"}}
	c.TaskEvents = []namedValue{{Value: "assignment_queued"}}
	c.AssignmentPayloadFields = []payloadField{{Name: "model_id"}}
	got := newEvidence(manifest{ID: "id", Contract: defaultContract}, c)
	if got.Status != "verified" || got.PayloadFieldCount != 1 {
		t.Fatalf("evidence = %#v", got)
	}
}
