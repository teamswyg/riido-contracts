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
	if blank(m.ID) || blank(m.RiidoTask) || blank(m.GeneratedDoc) || blank(m.Title) {
		return errors.New("id, riido_task, generated_doc, and title are required")
	}
	if len(m.Intro) == 0 || len(m.Sections) == 0 {
		return errors.New("intro and sections are required")
	}
	if blank(m.Workflow) || blank(m.EvidenceArtifact) {
		return errors.New("workflow and evidence_artifact are required")
	}
	if err := verifyLoop(m.Loop); err != nil {
		return err
	}
	return verifyPolicyShape(m)
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
