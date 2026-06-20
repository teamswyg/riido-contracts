package main

import "errors"

func verifyManifest(m manifest) error {
	if m.SchemaVersion == "" || m.ID == "" || m.Title == "" {
		return errors.New("manifest identity fields are required")
	}
	if m.GeneratedDoc == "" || m.Workflow == "" || m.EvidenceArtifact == "" {
		return errors.New("manifest artifact fields are required")
	}
	if m.DSLFixture == "" || m.IRFixture == "" || m.OpenAPIFixture == "" {
		return errors.New("manifest fixture paths are required")
	}
	if len(m.RequiredOperations) == 0 || len(m.Invariants) == 0 {
		return errors.New("manifest operations and invariants are required")
	}
	if len(m.FixtureRows) == 0 || len(m.NoDiffPathFragments) == 0 {
		return errors.New("manifest fixture rows and no-diff fragments are required")
	}
	return verifyLoop(m.Loop)
}
