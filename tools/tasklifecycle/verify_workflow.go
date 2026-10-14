package main

import (
	"errors"
	"os"
	"strings"
)

func verifyWorkflow(root string, m manifest) error {
	body, err := os.ReadFile(resolve(root, m.Workflow))
	if err != nil {
		return err
	}
	text := string(body)
	required := []string{"./tools/tasklifecycle", "-check-doc", "-evidence-out", m.EvidenceArtifact, "if-no-files-found: error"}
	for _, value := range required {
		if !strings.Contains(text, value) {
			return errors.New("workflow does not bind strict task lifecycle evidence")
		}
	}
	return nil
}
