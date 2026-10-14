package main

import (
	"bytes"
	"fmt"
	"os"
)

func verify() error {
	generated, err := generatedFixtures()
	if err != nil {
		return err
	}
	for path, want := range generated {
		got, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read %s: %w", path, err)
		}
		if !bytes.Equal(got, want) {
			return fmt.Errorf("%s drifted; run go run ./tools/apicontract generate", path)
		}
	}
	return nil
}
