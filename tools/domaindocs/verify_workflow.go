package main

import (
	"errors"
	"strings"
)

func verifyWorkflow(root string, m manifest) error {
	text, err := readRepoFile(root, m.Workflow)
	if err != nil {
		return err
	}
	required := []string{"./tools/domaindocs", "-check-doc", "-evidence-out", m.EvidenceArtifact, "if-no-files-found: error"}
	for _, value := range required {
		if !strings.Contains(text, value) {
			return errors.New("workflow does not bind strict domain docs evidence")
		}
	}
	return nil
}
