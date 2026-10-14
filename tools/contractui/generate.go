package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func generate() error {
	root, err := repoRoot()
	if err != nil {
		return err
	}
	data, err := buildBundle()
	if err != nil {
		return err
	}
	body, err := renderBundleJS(data)
	if err != nil {
		return err
	}
	path := filepath.Join(root, generatedPath)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(path, body, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", generatedPath, err)
	}
	return nil
}
