package main

import (
	"fmt"
	"os"
)

func generate() error {
	generated, err := generatedFixtures()
	if err != nil {
		return err
	}
	for path, body := range generated {
		if err := os.WriteFile(path, body, 0o644); err != nil {
			return fmt.Errorf("write %s: %w", path, err)
		}
	}
	return nil
}
