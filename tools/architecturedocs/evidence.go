package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type evidence struct {
	SchemaVersion       string       `json:"schema_version"`
	ID                  string       `json:"id"`
	Status              string       `json:"status"`
	PackageCount        int          `json:"package_count"`
	PublicGateCount     int          `json:"public_gate_count"`
	DownstreamGateCount int          `json:"downstream_gate_count"`
	RequiredDocCount    int          `json:"required_doc_count"`
	StaleScanPathCount  int          `json:"stale_scan_path_count"`
	CheckDoc            bool         `json:"check_doc"`
	EvidenceArtifact    string       `json:"evidence_artifact"`
	Loop                evidenceLoop `json:"loop"`
}

func newEvidence(m manifest, checkDoc bool) evidence {
	return evidence{
		SchemaVersion:       evidenceSchemaVersion,
		ID:                  m.ID,
		Status:              "verified",
		PackageCount:        len(m.Packages),
		PublicGateCount:     len(m.PublicGates),
		DownstreamGateCount: len(m.DownstreamGates),
		RequiredDocCount:    len(m.RequiredDocs),
		StaleScanPathCount:  len(m.StaleScanPaths),
		CheckDoc:            checkDoc,
		EvidenceArtifact:    "architecture-docs-evidence",
		Loop:                m.Loop,
	}
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
