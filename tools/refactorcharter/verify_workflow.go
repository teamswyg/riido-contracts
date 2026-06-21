package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func verifyWorkflowBinding(root string, c charter) error {
	body, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(c.Workflow)))
	if err != nil {
		return err
	}
	text := string(body)
	required := []string{
		"go run ./tools/refactorcharter -evidence-out out/refactorcharter-evidence.json",
		"name: " + c.EvidenceArtifact,
		"if-no-files-found: error",
	}
	for _, value := range required {
		if !strings.Contains(text, value) {
			return errors.New("workflow does not bind refactorcharter evidence")
		}
	}
	return nil
}
