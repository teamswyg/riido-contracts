package main

import (
	"errors"
	"os"
	"path/filepath"
)

func repoRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if hasRepoMarker(dir) {
			return dir, nil
		}
		next := filepath.Dir(dir)
		if next == dir {
			return "", errors.New("riido-contracts repo root not found")
		}
		dir = next
	}
}

func hasRepoMarker(dir string) bool {
	if _, err := os.Stat(filepath.Join(dir, "go.mod")); err != nil {
		return false
	}
	if _, err := os.Stat(filepath.Join(dir, "apicontract")); err != nil {
		return false
	}
	return true
}
