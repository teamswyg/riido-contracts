package main

import (
	"os"
	"path/filepath"
	"strings"
)

func manifestInventorySamples(root string, groups []manifestGroupCount, limit int) ([]manifestGroupSample, error) {
	byGroup := map[string][]string{}
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() && filepath.Base(path) == ".git" {
			return filepath.SkipDir
		}
		if entry.IsDir() || !strings.HasSuffix(path, ".riido.json") {
			return nil
		}
		group := manifestGroup(root, path)
		if len(byGroup[group]) < limit {
			byGroup[group] = append(byGroup[group], rel(root, path))
		}
		return nil
	})
	return orderedManifestSamples(groups, byGroup), err
}

func orderedManifestSamples(groups []manifestGroupCount, byGroup map[string][]string) []manifestGroupSample {
	samples := make([]manifestGroupSample, 0, len(groups))
	for _, group := range groups {
		samples = append(samples, manifestGroupSample{Group: group.Group, Paths: byGroup[group.Group]})
	}
	return samples
}
