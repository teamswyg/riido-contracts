package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const evidenceSchemaVersion = "riido-ssotdeps-evidence.v1"

type evidence struct {
	SchemaVersion           string `json:"schema_version"`
	ID                      string `json:"id"`
	Status                  string `json:"status"`
	FactsVerified           int    `json:"facts_verified"`
	RepoDependenciesChecked int    `json:"repo_dependencies_checked"`
	CheckDoc                bool   `json:"check_doc"`
}

func writeEvidence(path string, value evidence) error {
	body, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return fmt.Errorf("encode evidence: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("create evidence dir: %w", err)
	}
	if err := os.WriteFile(path, append(body, '\n'), 0o644); err != nil {
		return fmt.Errorf("write evidence: %w", err)
	}
	return nil
}
