package main

import "errors"

func verifyManifest(m manifest) error {
	if m.SchemaVersion == "" || m.ID == "" || m.Title == "" {
		return errors.New("manifest identity fields are required")
	}
	if m.GeneratedDoc == "" || m.Workflow == "" || m.EvidenceArtifact == "" {
		return errors.New("manifest artifact fields are required")
	}
	if m.Package == "" || len(m.Invariants) == 0 {
		return errors.New("manifest package and invariants are required")
	}
	return verifyLoop(m.Loop)
}
