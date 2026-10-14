package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func verifyNoStaleRuntimeWords(root string, m manifest) error {
	for _, rel := range m.StaleScanPaths {
		path, err := resolve(root, rel)
		if err != nil {
			return err
		}
		if err := scanNoStaleWords(path, m.StaleRuntimeWords); err != nil {
			return err
		}
	}
	return nil
}

func scanNoStaleWords(path string, words []string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return checkMarkdownFile(path, words)
	}
	return filepath.WalkDir(path, func(p string, entry os.DirEntry, err error) error {
		if err != nil || entry.IsDir() || filepath.Ext(p) != ".md" {
			return err
		}
		return checkMarkdownFile(p, words)
	})
}

func checkMarkdownFile(path string, words []string) error {
	body, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	text := string(body)
	for _, word := range words {
		if strings.Contains(text, word) {
			return fmt.Errorf("%s contains stale runtime/private wording %q", path, word)
		}
	}
	return nil
}
