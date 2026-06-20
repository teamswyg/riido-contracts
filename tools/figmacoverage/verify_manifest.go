package main

import (
	"errors"
	"fmt"
	"strings"
)

func verifyManifest(m manifest) error {
	if m.SchemaVersion != schemaVersion {
		return fmt.Errorf("schema_version = %q, want %q", m.SchemaVersion, schemaVersion)
	}
	if blank(m.ID) || blank(m.RiidoTask) || blank(m.HumanDoc) {
		return errors.New("id, riido_task, and human_doc are required")
	}
	if err := verifyFigmaSource(m.Figma); err != nil {
		return err
	}
	if err := verifyPages(m.ExpectedPages); err != nil {
		return err
	}
	if err := verifyEntries(m.ExpectedTopLevelNodes, m.Entries); err != nil {
		return err
	}
	if err := verifyAnnotationPolicy(m.APIAnnotationContentPolicy); err != nil {
		return err
	}
	if err := verifyAnnotationInventory(m.APIAnnotationContentPolicy, m.APIGeneratedAnnotationInventory); err != nil {
		return err
	}
	if err := verifyIncludeFiles(m); err != nil {
		return err
	}
	if len(m.VerifiedEvidenceNodes) == 0 {
		return errors.New("verified_evidence_nodes are required")
	}
	return nil
}

func verifyIncludeFiles(m manifest) error {
	fileLists := map[string][]string{
		"tool_limitation_files":                    m.ToolLimitationFiles,
		"expected_top_level_node_files":            m.ExpectedTopLevelNodeFiles,
		"page_inventory_files":                     m.PageInventoryFiles,
		"coverage_entry_files":                     m.CoverageEntryFiles,
		"non_ui_coverage_entry_files":              m.NonUICoverageEntryFiles,
		"api_generated_annotation_inventory_files": m.APIAnnotationInventoryFiles,
		"api_generated_annotation_files":           m.APIAnnotationFiles,
		"verified_evidence_node_files":             m.VerifiedEvidenceNodeFiles,
	}
	for name, files := range fileLists {
		if hasDuplicate(files) {
			return fmt.Errorf("%s has duplicate file refs", name)
		}
	}
	return nil
}

func hasDuplicate(values []string) bool {
	seen := map[string]bool{}
	for _, value := range values {
		if seen[value] {
			return true
		}
		seen[value] = true
	}
	return false
}

func blank(s string) bool {
	return strings.TrimSpace(s) == ""
}
