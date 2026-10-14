package main

import (
	"errors"
	"strings"
)

func verifyManifest(m manifest) error {
	if m.SchemaVersion != manifestSchema {
		return errors.New("unexpected schema_version")
	}
	if !filled(m.ID, m.Title, m.Summary, m.GeneratedDoc, m.Workflow, m.EvidenceArtifact) {
		return errors.New("id, title, summary, generated_doc, workflow, and evidence_artifact are required")
	}
	if m.Contract != defaultContract {
		return errors.New("manifest must point at assignment contract")
	}
	if len(m.Responsibilities) == 0 || len(m.Boundaries) == 0 || len(m.Invariants) == 0 {
		return errors.New("responsibilities, boundaries, and invariants are required")
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
