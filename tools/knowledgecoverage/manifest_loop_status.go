package main

import "path/filepath"

func manifestLoopStatus(root, path string) string {
	var doc map[string]any
	if !readManifestDoc(path, &doc) {
		return "missing"
	}
	if manifestDocHasLoop(doc) {
		return "direct"
	}
	if source, ok := doc["loop_source"].(string); ok && manifestSourceHasLoop(root, path, source) {
		return "delegated"
	}
	return "missing"
}

func manifestSourceHasLoop(root, path, source string) bool {
	if source == "" {
		return false
	}
	sourcePath := filepath.Join(root, filepath.FromSlash(source))
	if filepath.IsAbs(source) {
		sourcePath = source
	}
	if !fileWithinRoot(root, sourcePath) || sourcePath == path {
		return false
	}
	var doc map[string]any
	return readManifestDoc(sourcePath, &doc) && manifestDocHasLoop(doc)
}
