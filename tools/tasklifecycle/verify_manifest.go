package main

import (
	"errors"
	"strings"
)

func verifyManifest(m manifest) error {
	if m.SchemaVersion != manifestSchema {
		return errors.New("unexpected schema_version")
	}
	if !filled(m.ID, m.Title, m.Summary, m.GeneratedDoc, m.Workflow, m.EvidenceArtifact, m.SourceDSL) {
		return errors.New("id, title, summary, generated_doc, workflow, evidence_artifact, and source_dsl are required")
	}
	if len(m.Responsibilities) == 0 || len(m.Invariants) == 0 {
		return errors.New("responsibilities and invariants are required")
	}
	return verifyLoop(m.Loop)
}

func filled(values ...string) bool {
	for _, value := range values {
		if strings.TrimSpace(value) == "" {
			return false
		}
	}
	return true
}
