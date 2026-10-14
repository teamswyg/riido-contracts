package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func runFSMVerify(plan fsmRunPlan, out io.Writer) error {
	for _, file := range plan.Files {
		path := filepath.Join(plan.Root, filepath.FromSlash(file.Path))
		got, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read %s: %w", file.Path, err)
		}
		if !bytes.Equal(got, file.Body) {
			return fmt.Errorf("%s drifted; run go run ./tools/fsmgen generate", file.Path)
		}
	}
	if err := verifyReadmeSections(plan.Root, plan.Sections); err != nil {
		return err
	}
	fmt.Fprintf(out, "fsmgen: verified %d files and %d README sections\n", len(plan.Files), len(plan.Sections))
	return nil
}
