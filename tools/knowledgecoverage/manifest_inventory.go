package main

import (
	"sort"
	"strings"
)

func countManifests(root string) (int, []manifestGroupCount, error) {
	byGroup := map[string]int{}
	paths, err := manifestPaths(root)
	for _, path := range paths {
		byGroup[manifestGroup(root, path)]++
	}
	return len(paths), manifestGroups(byGroup), err
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
