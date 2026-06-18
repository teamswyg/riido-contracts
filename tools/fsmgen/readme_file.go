package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func writeReadmeSections(root string, sections []readmeSection) error {
	path := filepath.Join(root, readmePath)
	body, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s: %w", readmePath, err)
	}
	updated, err := replaceReadmeSections(string(body), sections)
	if err != nil {
		return err
	}
	if err := os.WriteFile(path, []byte(updated), 0o644); err != nil {
		return fmt.Errorf("write %s: %w", readmePath, err)
	}
	return nil
}

func verifyReadmeSections(root string, sections []readmeSection) error {
	path := filepath.Join(root, readmePath)
	body, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s: %w", readmePath, err)
	}
	return verifyReadmeSectionContent(string(body), sections)
}
