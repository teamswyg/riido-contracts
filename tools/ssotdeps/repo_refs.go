package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func readLocalRef(root, path string) (string, error) {
	if filepath.IsAbs(path) || strings.Contains(path, "..") {
		return "", fmt.Errorf("path %q must be repo-relative", path)
	}
	body, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(path)))
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func findRepoRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if isRepoRoot(dir) {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", errors.New("could not find repository root")
		}
		dir = parent
	}
}

func isRepoRoot(dir string) bool {
	if _, err := os.Stat(filepath.Join(dir, "go.mod")); err != nil {
		return false
	}
	manifestPath := filepath.Join(dir, filepath.FromSlash(defaultManifest))
	_, err := os.Stat(manifestPath)
	return err == nil
}
