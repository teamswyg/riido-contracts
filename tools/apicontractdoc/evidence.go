package main

import "encoding/json"

type evidence struct {
	SchemaVersion          string           `json:"schema_version"`
	ID                     string           `json:"id"`
	Status                 string           `json:"status"`
	FixtureCount           int              `json:"fixture_count"`
	OperationCount         int              `json:"operation_count"`
	GeneratedPathCount     int              `json:"generated_path_count"`
	RequiredGeneratedPaths []string         `json:"required_generated_paths"`
	EvidenceArtifact       string           `json:"evidence_artifact"`
	Workflow               string           `json:"workflow"`
	Fixtures               []fixtureSummary `json:"fixtures"`
	Loop                   evidenceLoop     `json:"loop"`
}

func writeEvidence(path string, m manifest, summaries []fixtureSummary) error {
	body, err := json.MarshalIndent(newEvidence(m, summaries), "", "  ")
	if err != nil {
		return err
	}
	return writeFile(path, append(body, '\n'))
}
