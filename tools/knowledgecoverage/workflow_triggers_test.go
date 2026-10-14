package main

import (
	"strings"
	"testing"
)

func TestVerifyWorkflowTriggerPathsRejectsMissingPullRequestPath(t *testing.T) {
	text := `on:
  push:
    paths:
      - "docs/**"
  pull_request:
    paths:
      - "README.md"
`
	err := verifyWorkflowTriggerPaths(text, []string{"docs/**"})
	if err == nil {
		t.Fatal("expected missing pull_request path")
	}
	if !strings.Contains(err.Error(), `pull_request.paths missing "docs/**"`) {
		t.Fatalf("error = %v", err)
	}
}

func TestVerifyWorkflowTriggerPathsAcceptsQuotedAndBarePaths(t *testing.T) {
	text := `on:
  push:
    branches: [main]
    paths:
      - "docs/**"
      - go.mod
  pull_request:
    paths:
      - 'docs/**'
      - go.mod
`
	if err := verifyWorkflowTriggerPaths(text, []string{"docs/**", "go.mod"}); err != nil {
		t.Fatalf("verifyWorkflowTriggerPaths() error = %v", err)
	}
}
