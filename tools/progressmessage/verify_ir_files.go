package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/teamswyg/riido-contracts/progressmessage"
)

func verifyGeneratedIR(root string, ir progressmessage.IRDocument) error {
	files, err := generatedIRFilesFor(ir)
	if err != nil {
		return err
	}
	if err := verifyFileBytes(root, irPath, files.Root); err != nil {
		return err
	}
	expected := make([]string, 0, len(files.Messages))
	for path, want := range files.Messages {
		expected = append(expected, path)
		if err := verifyFileBytes(root, path, want); err != nil {
			return err
		}
	}
	return verifyNoStaleIRMessages(root, expected)
}

func verifyFileBytes(root, path string, want []byte) error {
	got, err := os.ReadFile(resolve(root, path))
	if err != nil {
		return fmt.Errorf("read %s: %w", path, err)
	}
	if !bytes.Equal(got, want) {
		return fmt.Errorf("%s drifted; run go run ./tools/progressmessage generate", path)
	}
	return nil
}

func verifyNoStaleIRMessages(root string, expected []string) error {
	actual, err := filepath.Glob(resolve(root, filepath.Join(irMessageDir, "*.riido.json")))
	if err != nil {
		return err
	}
	sort.Strings(actual)
	sort.Strings(expected)
	if len(actual) != len(expected) {
		return fmt.Errorf("%s has %d files, want %d", irMessageDir, len(actual), len(expected))
	}
	for i, path := range expected {
		if actual[i] != resolve(root, path) {
			return fmt.Errorf("%s has stale or missing generated files", irMessageDir)
		}
	}
	return nil
}
