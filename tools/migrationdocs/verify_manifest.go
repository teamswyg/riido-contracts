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
	if blank(m.ID) || blank(m.RiidoTask) || blank(m.GeneratedDoc) || blank(m.Goal) {
		return errors.New("id, riido_task, generated_doc, and goal are required")
	}
	if len(m.PromotionRule.Conditions) == 0 || blank(m.PromotionRule.Fallback) {
		return errors.New("promotion rule is required")
	}
	if len(m.CandidateContracts) == 0 || len(m.MigrationSlices) == 0 {
		return errors.New("candidate contracts and migration slices are required")
	}
	if blank(m.Workflow) || blank(m.EvidenceArtifact) {
		return errors.New("workflow and evidence_artifact are required")
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
