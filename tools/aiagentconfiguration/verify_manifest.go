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
	if len(m.RequiredOperations) == 0 || len(m.RequiredSchemaFields) == 0 {
		return errors.New("manifest operation and schema fields are required")
	}
	if len(m.RequiredPolicies) == 0 || len(m.RequiredEnums) == 0 {
		return errors.New("manifest policy and enum fields are required")
	}
	return verifyLoop(m.Loop)
}
