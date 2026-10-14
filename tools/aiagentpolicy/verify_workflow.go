package main

import (
	"fmt"
	"os"
	"strings"
)

func verifyWorkflow(root string, m manifest) error {
	body, err := os.ReadFile(resolve(root, m.Workflow))
	if err != nil {
		return err
	}
	workflow := string(body)
	required := []string{
		"./tools/aiagentpolicy",
		"-check-doc",
		"-evidence-out",
		m.EvidenceArtifact,
		"if-no-files-found: error",
	}
	for _, item := range required {
		if !strings.Contains(workflow, item) {
			return fmt.Errorf("%s missing workflow evidence item %q", m.Workflow, item)
		}
	}
	return nil
}
