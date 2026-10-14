package main

import "strings"

type manifestLoopSource struct {
	ID           string   `json:"id"`
	LoopSource   string   `json:"loop_source"`
	Paths        []string `json:"paths"`
	PathPrefixes []string `json:"path_prefixes"`
}

func externalManifestSourceHasLoop(root, path string, sources []manifestLoopSource) bool {
	for _, source := range sources {
		if !source.matches(rel(root, path)) {
			continue
		}
		if manifestSourceHasLoop(root, path, source.LoopSource) {
			return true
		}
	}
	return false
}

func (source manifestLoopSource) matches(path string) bool {
	for _, exact := range source.Paths {
		if path == exact {
			return true
		}
	}
	for _, prefix := range source.PathPrefixes {
		if strings.HasPrefix(path, prefix) {
			return true
		}
	}
	return false
}
