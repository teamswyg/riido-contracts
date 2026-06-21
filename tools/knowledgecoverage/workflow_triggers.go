package main

import (
	"fmt"
	"strings"
)

func verifyWorkflowTriggerPaths(text string, required []string) error {
	for _, event := range []string{"push", "pull_request"} {
		paths := workflowEventPaths(text, event)
		for _, path := range required {
			if !paths[path] {
				return fmt.Errorf("workflow %s.paths missing %q", event, path)
			}
		}
	}
	return nil
}

func workflowEventPaths(text, event string) map[string]bool {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		if strings.TrimSpace(line) == event+":" {
			return workflowPathsInBlock(lines, i+1, leadingSpaces(line))
		}
	}
	return map[string]bool{}
}

func workflowPathsInBlock(lines []string, start, eventIndent int) map[string]bool {
	out := map[string]bool{}
	for i := start; i < len(lines); i++ {
		line := lines[i]
		if strings.TrimSpace(line) == "" {
			continue
		}
		if leadingSpaces(line) <= eventIndent {
			break
		}
		if strings.TrimSpace(line) == "paths:" {
			return workflowPathItems(lines, i+1, leadingSpaces(line))
		}
	}
	return out
}

func workflowPathItems(lines []string, start, pathsIndent int) map[string]bool {
	out := map[string]bool{}
	for i := start; i < len(lines); i++ {
		line := lines[i]
		if strings.TrimSpace(line) == "" {
			continue
		}
		if leadingSpaces(line) <= pathsIndent {
			break
		}
		if value, ok := strings.CutPrefix(strings.TrimSpace(line), "- "); ok {
			out[strings.Trim(value, `"'`)] = true
		}
	}
	return out
}

func leadingSpaces(line string) int {
	return len(line) - len(strings.TrimLeft(line, " "))
}
