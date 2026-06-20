package main

import (
	"os"
	"path/filepath"
)

func writeRepoFile(root, rel, body string) error {
	path, err := resolve(root, rel)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, []byte(body), 0o644)
}
