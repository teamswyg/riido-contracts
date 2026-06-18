package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func runEnumGenerate(plan enumRunPlan, out io.Writer) error {
	for name, body := range plan.Files {
		path := filepath.Join(plan.Root, filepath.FromSlash(name))
		if err := os.WriteFile(path, body, 0o644); err != nil {
			return fmt.Errorf("write %s: %w", name, err)
		}
	}
	fmt.Fprintf(out, "enumgen: generated %d files\n", len(plan.Files))
	return nil
}
