package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func runFSMGenerate(plan fsmRunPlan, out io.Writer) error {
	for _, file := range plan.Files {
		path := filepath.Join(plan.Root, filepath.FromSlash(file.Path))
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			return fmt.Errorf("mkdir %s: %w", filepath.Dir(file.Path), err)
		}
		if err := os.WriteFile(path, file.Body, 0o644); err != nil {
			return fmt.Errorf("write %s: %w", file.Path, err)
		}
	}
	if err := writeReadmeSections(plan.Root, plan.Sections); err != nil {
		return err
	}
	fmt.Fprintf(out, "fsmgen: generated %d files and %d README sections\n", len(plan.Files), len(plan.Sections))
	return nil
}
