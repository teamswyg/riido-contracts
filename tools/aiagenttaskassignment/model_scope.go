package main

import "strings"

func noDiffPathsAbsent(dsl contractFixture, fragments []string) bool {
	for _, op := range dsl.Operations {
		if !looksLikeAgentAssignment(op) {
			continue
		}
		for _, fragment := range fragments {
			if strings.Contains(op.Path, fragment) {
				return false
			}
		}
	}
	return true
}

func looksLikeAgentAssignment(op operation) bool {
	if !strings.Contains(op.Path, "/ai-agent/") {
		return false
	}
	id := strings.ToLower(op.OperationID)
	return strings.Contains(id, "assign") ||
		strings.Contains(id, "thread") ||
		strings.Contains(id, "stop")
}
