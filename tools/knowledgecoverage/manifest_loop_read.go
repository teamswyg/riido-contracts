package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

func readManifestDoc(path string, v any) bool {
	body, err := os.ReadFile(path)
	if err != nil {
		return false
	}
	return json.Unmarshal(body, v) == nil
}

func fileWithinRoot(root, path string) bool {
	rel, err := filepath.Rel(root, path)
	if err != nil || filepath.IsAbs(rel) {
		return false
	}
	return rel != ".." && !strings.HasPrefix(rel, ".."+string(filepath.Separator))
}
