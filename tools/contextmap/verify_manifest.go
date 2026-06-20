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
	if blank(m.ID) || blank(m.RiidoTask) || blank(m.GeneratedDoc) || blank(m.Module) {
		return errors.New("id, riido_task, generated_doc, and module are required")
	}
	if blank(m.Role) || blank(m.BoundaryRule) || blank(m.Workflow) || blank(m.EvidenceArtifact) {
		return errors.New("role, boundary_rule, workflow, and evidence_artifact are required")
	}
	if len(m.OwnedContexts) == 0 || len(m.NonOwnedContexts) == 0 {
		return errors.New("owned and non-owned contexts are required")
	}
	if len(m.DirectionRules) == 0 || len(m.SSOTLinks) == 0 {
		return errors.New("direction rules and ssot links are required")
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
