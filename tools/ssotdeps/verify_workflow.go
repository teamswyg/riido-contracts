package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func verifyWorkflowBinding(root string, m manifest) error {
	body, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(m.Workflow)))
	if err != nil {
		return err
	}
	text := string(body)
	required := []string{
		"go run ./tools/ssotdeps verify -check-doc -evidence-out out/ssotdeps-evidence.json",
		"name: " + m.EvidenceArtifact,
		"if-no-files-found: error",
	}
	for _, value := range required {
		if !strings.Contains(text, value) {
			return errors.New("workflow does not bind ssotdeps evidence")
		}
	}
	return nil
}
