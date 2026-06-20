package main

import "strings"

func sameOperationTuple(left, right operation) bool {
	return left.OperationID != "" &&
		left.OperationID == right.OperationID &&
		left.Kind == right.Kind &&
		left.Method == right.Method &&
		left.Path == right.Path
}

func openAPIMatches(openapi openAPIDoc, op operation) bool {
	methods, ok := openapi.Paths[op.Path]
	if !ok {
		return false
	}
	projected, ok := methods[strings.ToLower(op.Method)]
	return ok && projected.OperationID == op.OperationID
}

func isOnboardingOperation(op operation) bool {
	return strings.Contains(op.Path, "/onboarding/fixtures")
}
