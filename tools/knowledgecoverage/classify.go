package main

import (
	"os"
	"strings"
)

func classify(record docRecord) string {
	if record.HasGeneratedMarker {
		return "generated_reader"
	}
	if record.HasExecutableMarker || record.HasAdjacentManifest {
		return "executable_reader"
	}
	return "manual_reader"
}

func containsAny(text string, needles []string) bool {
	for _, needle := range needles {
		if strings.Contains(text, needle) {
			return true
		}
	}
	return false
}

func adjacentManifestExists(path string) bool {
	manifestPath := strings.TrimSuffix(path, ".md") + ".riido.json"
	_, err := os.Stat(manifestPath)
	return err == nil
}
