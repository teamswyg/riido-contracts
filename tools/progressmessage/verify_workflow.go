package main

import (
	"errors"
	"os"
	"strings"
)

func verifyWorkflow(root string, m docManifest) error {
	body, err := os.ReadFile(resolve(root, m.Workflow))
	if err != nil {
		return err
	}
	text := string(body)
	required := []string{"./tools/progressmessage", "-check-doc", "-evidence-out", m.EvidenceArtifact, "if-no-files-found: error"}
	for _, value := range required {
		if !strings.Contains(text, value) {
			return errors.New("workflow does not bind strict progress message evidence")
		}
	}
	return nil
}
