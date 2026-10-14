package main

import "strings"

func noDiffPathsClean(openapi openAPIDoc, fragments []string) bool {
	for path := range openapi.Paths {
		for _, fragment := range fragments {
			if strings.Contains(strings.ToLower(path), strings.ToLower(fragment)) {
				return false
			}
		}
	}
	return true
}
