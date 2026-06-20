package main

import (
	"errors"
	"os"
	"strings"
)

func verifyManifest(root string, m manifest) error {
	if m.SchemaVersion != schemaVersion {
		return errors.New("unexpected schema_version")
	}
	if !filled(m.ID, m.Title, m.GeneratedDoc, m.Workflow, m.EvidenceArtifact) {
		return errors.New("id, title, generated_doc, workflow, and evidence_artifact are required")
	}
	if len(m.ScanRoots) == 0 || len(m.GeneratedMarkers) == 0 {
		return errors.New("scan_roots and generated_markers are required")
	}
	if !filled(m.Loop.Observation, m.Loop.Hypothesis, m.Loop.Execute, m.Loop.Evaluate, m.Loop.Retrospective) {
		return errors.New("complete evidence loop is required")
	}
	return verifyWorkflow(root, m)
}

func filled(values ...string) bool {
	for _, value := range values {
		if strings.TrimSpace(value) == "" {
			return false
		}
	}
	return true
}

func verifyWorkflow(root string, m manifest) error {
	body, err := os.ReadFile(resolve(root, m.Workflow))
	if err != nil {
		return err
	}
	text := string(body)
	required := []string{"./tools/knowledgecoverage", "-check-doc", "-evidence-out", m.EvidenceArtifact, "if-no-files-found: error"}
	for _, value := range required {
		if !strings.Contains(text, value) {
			return errors.New("workflow does not bind strict knowledge coverage evidence")
		}
	}
	return nil
}
