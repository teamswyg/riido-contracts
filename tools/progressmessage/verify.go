package main

import (
	"bytes"
	"fmt"
	"os"
)

func verify() error {
	want, err := generatedIR()
	if err != nil {
		return err
	}
	got, err := os.ReadFile(irPath)
	if err != nil {
		return fmt.Errorf("read %s: %w", irPath, err)
	}
	if !bytes.Equal(got, want) {
		return fmt.Errorf("%s drifted; run go run ./tools/progressmessage generate", irPath)
	}
	return nil
}
