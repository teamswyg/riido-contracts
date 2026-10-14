package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func generatedFile(root, rel string, cfg scanConfig) bool {
	path := "/" + rel
	for _, fragment := range cfg.GeneratedPathFragments {
		if strings.Contains(path, fragment) {
			return true
		}
	}
	return hasGeneratedMarker(filepath.Join(root, filepath.FromSlash(rel)), cfg.GeneratedMarkers)
}

func hasGeneratedMarker(path string, markers []string) bool {
	if len(markers) == 0 {
		return false
	}
	file, err := os.Open(path)
	if err != nil {
		return false
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for i := 0; i < 8 && scanner.Scan(); i++ {
		line := scanner.Text()
		for _, marker := range markers {
			if strings.Contains(line, marker) {
				return true
			}
		}
	}
	return false
}
