package main

import (
	"errors"
	"strings"
)

func verifyManifest(root string, m docManifest) error {
	if m.SchemaVersion != docSchemaVersion {
		return errors.New("unexpected schema_version")
	}
	if !filled(m.ID, m.Title, m.Summary, m.GeneratedDoc, m.Workflow, m.EvidenceArtifact) {
		return errors.New("id, title, summary, generated_doc, workflow, and evidence_artifact are required")
	}
	if m.DSL != dslPath || m.IR != irPath {
		return errors.New("manifest must point at the progressmessage DSL and IR")
	}
	if len(m.Rules) == 0 || len(m.Projection) == 0 {
		return errors.New("rules and projection are required")
	}
	if !filled(m.Loop.Observation, m.Loop.Hypothesis, m.Loop.Execute, m.Loop.Evaluate, m.Loop.Retrospective) {
		return errors.New("complete evidence loop is required")
	}
	return verifyWorkflow(root, m)
}

func filled(values ...string) bool {
	for _, value := range values {
		if strings.TrimSpace(value) == "" {
			return false
		}
	}
	return true
}
