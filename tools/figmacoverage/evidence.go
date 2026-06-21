package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const evidenceSchemaVersion = "riido-figma-coverage-evidence.v1"

type evidence struct {
	SchemaVersion                     string       `json:"schema_version"`
	ID                                string       `json:"id"`
	Status                            string       `json:"status"`
	Workflow                          string       `json:"workflow"`
	EvidenceArtifact                  string       `json:"evidence_artifact"`
	ToolLimitationFilesLoaded         int          `json:"tool_limitation_files_loaded"`
	ExpectedTopLevelNodeFilesLoaded   int          `json:"expected_top_level_node_files_loaded"`
	PageInventoryFilesLoaded          int          `json:"page_inventory_files_loaded"`
	CoverageEntryFilesLoaded          int          `json:"coverage_entry_files_loaded"`
	NonUICoverageEntryFilesLoaded     int          `json:"non_ui_coverage_entry_files_loaded"`
	APIAnnotationInventoryFilesLoaded int          `json:"api_annotation_inventory_files_loaded"`
	APIAnnotationFilesLoaded          int          `json:"api_annotation_files_loaded"`
	VerifiedEvidenceNodeFilesLoaded   int          `json:"verified_evidence_node_files_loaded"`
	EntriesVerified                   int          `json:"entries_verified"`
	GeneratedAnnotationsChecked       int          `json:"generated_annotations_checked"`
	EvidenceNodesVerified             int          `json:"evidence_nodes_verified"`
	CheckDoc                          bool         `json:"check_doc"`
	Loop                              evidenceLoop `json:"loop"`
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
