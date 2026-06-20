package main

import (
	"errors"
	"os"
	"path/filepath"
)

func resolveRoot(root string) (string, error) {
	if root != "." {
		return filepath.Abs(root)
	}
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for dir := wd; ; dir = filepath.Dir(dir) {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		if next := filepath.Dir(dir); next == dir {
			return "", errors.New("repository root not found")
		}
	}
}

func resolve(root, path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(root, filepath.FromSlash(path))
}
