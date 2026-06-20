package main

import "errors"

func verifyManifest(m manifest) error {
	if m.SchemaVersion == "" || m.ID == "" || m.Title == "" {
		return errors.New("manifest identity fields are required")
	}
	if m.GeneratedDoc == "" || m.Workflow == "" || m.EvidenceArtifact == "" {
		return errors.New("manifest artifact fields are required")
	}
	if m.ContractID == "" || m.DSLFixture == "" || m.IRFixture == "" || m.OpenAPIFixture == "" {
		return errors.New("manifest fixture anchors are required")
	}
	if len(m.RequiredStreamVariants) == 0 || len(m.CodegenRules) == 0 || len(m.Invariants) == 0 {
		return errors.New("manifest rules, invariants, and stream variants are required")
	}
	return verifyLoop(m.Loop)
}
