package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func resolveRoot(root string) (string, error) {
	if root == "" {
		root = "."
	}
	current, err := filepath.Abs(root)
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(current, "go.mod")); err == nil {
			return current, nil
		}
		next := filepath.Dir(current)
		if next == current {
			return "", fmt.Errorf("go.mod not found above %s", root)
		}
		current = next
	}
}
