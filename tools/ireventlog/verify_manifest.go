package main

import "errors"

func verifyManifest(m manifest) error {
	if m.SchemaVersion == "" || m.ID == "" || m.Title == "" {
		return errors.New("manifest identity fields are required")
	}
	if m.GeneratedDoc == "" || m.Workflow == "" || m.EvidenceArtifact == "" {
		return errors.New("manifest artifact paths are required")
	}
	if m.SourceDSL == "" || m.Package == "" {
		return errors.New("manifest executable source fields are required")
	}
	if len(m.Invariants) == 0 || len(m.Responsibilities) == 0 {
		return errors.New("manifest invariants and responsibilities are required")
	}
	return verifyLoop(m.Loop)
}
