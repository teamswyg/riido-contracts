package main

import (
	"slices"
	"strings"
)

func cell(s string) string {
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "|", "\\|")
	return s
}

func countStatuses(entries []coverageEntry) map[string]int {
	counts := map[string]int{}
	for _, entry := range entries {
		counts[entry.CoverageStatus]++
	}
	return counts
}

func sortedKeys(values map[string]int) []string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	return keys
}
