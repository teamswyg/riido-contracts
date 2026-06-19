package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const evidenceSchemaVersion = "riido-refactor-charter-evidence.v1"

type evidence struct {
	SchemaVersion  string    `json:"schema_version"`
	ID             string    `json:"id"`
	Status         string    `json:"status"`
	Mode           string    `json:"mode"`
	FilesScanned   int       `json:"files_scanned"`
	OverTarget     int       `json:"over_target"`
	TargetMaxLines int       `json:"target_max_lines"`
	Findings       []finding `json:"findings"`
}

func newEvidence(c charter, report scanReport) evidence {
	return evidence{
		SchemaVersion:  evidenceSchemaVersion,
		ID:             c.ID,
		Status:         "verified",
		Mode:           c.Mode,
		FilesScanned:   report.FilesScanned,
		OverTarget:     len(report.Findings),
		TargetMaxLines: c.LineBudget.TargetMaxLines,
		Findings:       report.Findings,
	}
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
