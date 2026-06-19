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
	if blank(m.ID) || blank(m.RiidoTask) || blank(m.HumanDoc) {
		return errors.New("id, riido_task, and human_doc are required")
	}
	if err := verifyFigmaSource(m.Figma); err != nil {
		return err
	}
	if err := verifyPages(m.ExpectedPages); err != nil {
		return err
	}
	if err := verifyEntries(m.ExpectedTopLevelNodes, m.Entries); err != nil {
		return err
	}
	if err := verifyAnnotationPolicy(m.APIAnnotationContentPolicy); err != nil {
		return err
	}
	if err := verifyAnnotationInventory(m.APIAnnotationContentPolicy, m.APIGeneratedAnnotationInventory); err != nil {
		return err
	}
	if len(m.VerifiedEvidenceNodes) == 0 {
		return errors.New("verified_evidence_nodes are required")
	}
	return nil
}

func blank(s string) bool {
	return strings.TrimSpace(s) == ""
}
