package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func runEnumVerify(plan enumRunPlan, out io.Writer) error {
	for name, want := range plan.Files {
		path := filepath.Join(plan.Root, filepath.FromSlash(name))
		got, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read %s: %w", name, err)
		}
		if !bytes.Equal(got, want) {
			return fmt.Errorf("%s drifted; run go run ./tools/enumgen generate", name)
		}
	}
	fmt.Fprintf(out, "enumgen: verified %d files\n", len(plan.Files))
	return nil
}
