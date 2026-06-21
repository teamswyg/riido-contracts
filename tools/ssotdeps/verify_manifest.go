package main

import (
	"errors"
	"fmt"
	"strings"
)

func verifyManifest(m manifest, root string) error {
	if m.SchemaVersion != schemaVersion {
		return fmt.Errorf("schema_version = %q, want %q", m.SchemaVersion, schemaVersion)
	}
	if err := requireID("manifest id", m.ID); err != nil {
		return err
	}
	if strings.TrimSpace(m.RiidoTask) == "" {
		return errors.New("riido_task is required")
	}
	if strings.TrimSpace(m.Workflow) == "" || strings.TrimSpace(m.EvidenceArtifact) == "" {
		return errors.New("workflow and evidence_artifact are required")
	}
	if !completeLoop(m.Loop) {
		return errors.New("complete evidence loop is required")
	}
	humanDoc, err := readLocalRef(root, m.HumanDoc)
	if err != nil {
		return fmt.Errorf("human_doc: %w", err)
	}
	if len(m.Facts) == 0 {
		return errors.New("facts are required")
	}
	if !stringsSorted(m.FactFiles) {
		return errors.New("fact_files must be sorted")
	}
	if !stringsSorted(m.RepoDependencyFiles) {
		return errors.New("repo_dependency_files must be sorted")
	}
	if !factsSorted(m.Facts) {
		return errors.New("facts must be sorted by id")
	}
	factIDs := map[string]bool{}
	for _, f := range m.Facts {
		if err := verifyFact(root, humanDoc, f); err != nil {
			return fmt.Errorf("fact %q: %w", f.ID, err)
		}
		if factIDs[f.ID] {
			return fmt.Errorf("duplicate fact id %q", f.ID)
		}
		factIDs[f.ID] = true
	}
	if !repoDependenciesSorted(m.RepoDependencies) {
		return errors.New("repo_dependencies must be sorted by id")
	}
	if err := verifyRepoDependencies(m.RepoDependencies, factIDs); err != nil {
		return err
	}
	return nil
}
