package main

import (
	"errors"
	"fmt"
	"strings"
)

func verifyManifest(m manifest) error {
	if m.SchemaVersion != schemaVersion {
		return fmt.Errorf("schema_version = %q, want %q", m.SchemaVersion, schemaVersion)
	}
	if blank(m.ID) || blank(m.RiidoTask) || blank(m.GeneratedDoc) || blank(m.Summary) {
		return errors.New("id, riido_task, generated_doc, and summary are required")
	}
	if blank(m.Workflow) || blank(m.EvidenceArtifact) {
		return errors.New("workflow and evidence_artifact are required")
	}
	if len(m.ArchitectureLinks) == 0 || len(m.Changes) == 0 {
		return errors.New("architecture_links and changes are required")
	}
	if len(m.ExternalBoundaries) == 0 || blank(m.OpenQuestions.Path) {
		return errors.New("external_boundaries and open_questions are required")
	}
	return verifyLoop(m.Loop)
}

func verifyLoop(loop evidenceLoop) error {
	if blank(loop.Observation) || blank(loop.Hypothesis) || blank(loop.Execute) {
		return errors.New("loop observation, hypothesis, and execute are required")
	}
	if blank(loop.Evaluate) || blank(loop.Retrospective) {
		return errors.New("loop evaluate and retrospective are required")
	}
	return nil
}

func blank(value string) bool {
	return strings.TrimSpace(value) == ""
}
