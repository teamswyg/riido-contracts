package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type evidence struct {
	SchemaVersion          string       `json:"schema_version"`
	ID                     string       `json:"id"`
	Status                 string       `json:"status"`
	CandidateCount         int          `json:"candidate_count"`
	SliceCount             int          `json:"slice_count"`
	SliceDoesCount         int          `json:"slice_does_count"`
	ValidationCommandCount int          `json:"validation_command_count"`
	WorkMapCount           int          `json:"work_map_count"`
	CheckDoc               bool         `json:"check_doc"`
	EvidenceArtifact       string       `json:"evidence_artifact"`
	Loop                   evidenceLoop `json:"loop"`
}

func newEvidence(m manifest, checkDoc bool) evidence {
	return evidence{
		SchemaVersion:          evidenceSchemaVersion,
		ID:                     m.ID,
		Status:                 "verified",
		CandidateCount:         len(m.CandidateContracts),
		SliceCount:             len(m.MigrationSlices),
		SliceDoesCount:         sliceDoesCount(m),
		ValidationCommandCount: len(m.ValidationGates.RequiredCommands) + len(m.ValidationGates.ArchitectureCommands),
		WorkMapCount:           len(m.MigrationWorkMap),
		CheckDoc:               checkDoc,
		EvidenceArtifact:       m.EvidenceArtifact,
		Loop:                   m.Loop,
	}
}

func sliceDoesCount(m manifest) int {
	total := 0
	for _, slice := range m.MigrationSlices {
		total += len(slice.Does)
	}
	return total
}

func writeEvidence(path string, value evidence) error {
	body, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, append(body, '\n'), 0o644)
}
