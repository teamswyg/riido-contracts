package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func loadFSMMetadata(path string) (map[string]fsmMetadata, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("resolve %s: %w", path, err)
	}
	return loadFSMMetadataAt(abs, filepath.Dir(abs), map[string]bool{})
}

func loadFSMMetadataAt(path, base string, seen map[string]bool) (map[string]fsmMetadata, error) {
	if seen[path] {
		return nil, fmt.Errorf("include cycle at %s", path)
	}
	seen[path] = true
	defer delete(seen, path)
	body, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}
	root, err := parseSExpr(string(body))
	if err != nil {
		return nil, err
	}
	return fsmMetadataFromLoadedNode(root, base, seen)
}
