package main

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func countManifests(root string) (int, []manifestGroupCount, error) {
	count := 0
	byGroup := map[string]int{}
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() && filepath.Base(path) == ".git" {
			return filepath.SkipDir
		}
		if entry.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".riido.json") {
			count++
			byGroup[manifestGroup(root, path)]++
		}
		return nil
	})
	return count, manifestGroups(byGroup), err
}

func manifestGroup(root, path string) string {
	parts := strings.Split(rel(root, path), "/")
	if len(parts) == 1 {
		return "."
	}
	return parts[0]
}

func manifestGroups(byGroup map[string]int) []manifestGroupCount {
	var groups []manifestGroupCount
	for group, count := range byGroup {
		groups = append(groups, manifestGroupCount{Group: group, Count: count})
	}
	sort.Slice(groups, func(i, j int) bool {
		if groups[i].Count == groups[j].Count {
			return groups[i].Group < groups[j].Group
		}
		return groups[i].Count > groups[j].Count
	})
	return groups
}
