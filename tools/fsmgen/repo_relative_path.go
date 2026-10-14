package main

import (
	"errors"
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

func cleanRepoRelativePath(value string) (string, error) {
	normalized := filepath.ToSlash(value)
	if normalized == "" {
		return "", errors.New("path is empty")
	}
	if strings.ContainsRune(normalized, 0) {
		return "", errors.New("path contains NUL")
	}
	if strings.HasPrefix(normalized, "/") {
		return "", fmt.Errorf("path %s must be repository-relative", value)
	}
	clean := path.Clean(normalized)
	if clean == "." || clean == ".." || strings.HasPrefix(clean, "../") {
		return "", fmt.Errorf("path %s must stay inside repository", value)
	}
	return clean, nil
}

func patternSources(metadata map[string]fsmMetadata) (map[string]bool, error) {
	sources := map[string]bool{}
	for _, spec := range metadata {
		source, err := cleanRepoRelativePath(spec.PatternSource)
		if err != nil {
			return nil, fmt.Errorf("transitions %s pattern-source: %w", spec.TransitionName, err)
		}
		sources[source] = true
	}
	return sources, nil
}
