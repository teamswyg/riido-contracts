package main

import (
	"bytes"
	"fmt"
	"os"
)

func verify() error {
	root, err := repoRoot()
	if err != nil {
		return err
	}
	data, err := buildBundle()
	if err != nil {
		return err
	}
	want, err := renderBundleJS(data)
	if err != nil {
		return err
	}
	got, err := os.ReadFile(root + "/" + generatedPath)
	if err != nil {
		return fmt.Errorf("read %s: %w", generatedPath, err)
	}
	if !bytes.Equal(got, want) {
		return fmt.Errorf("%s drifted; run go run ./tools/contractui generate", generatedPath)
	}
	return nil
}
