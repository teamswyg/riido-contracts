package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type evidence struct {
	SchemaVersion         string       `json:"schema_version"`
	ID                    string       `json:"id"`
	Status                string       `json:"status"`
	ChangeCount           int          `json:"change_count"`
	ChangeItemCount       int          `json:"change_item_count"`
	ArchitectureLinkCount int          `json:"architecture_link_count"`
	ExternalBoundaryCount int          `json:"external_boundary_count"`
	CheckDoc              bool         `json:"check_doc"`
	EvidenceArtifact      string       `json:"evidence_artifact"`
	Loop                  evidenceLoop `json:"loop"`
}

func newEvidence(m manifest, checkDoc bool) evidence {
	return evidence{
		SchemaVersion:         evidenceSchemaVersion,
		ID:                    m.ID,
		Status:                "verified",
		ChangeCount:           len(m.Changes),
		ChangeItemCount:       changeItemCount(m),
		ArchitectureLinkCount: len(m.ArchitectureLinks),
		ExternalBoundaryCount: len(m.ExternalBoundaries),
		CheckDoc:              checkDoc,
		EvidenceArtifact:      m.EvidenceArtifact,
		Loop:                  m.Loop,
	}
}

func changeItemCount(m manifest) int {
	total := 0
	for _, change := range m.Changes {
		total += len(change.Items)
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
