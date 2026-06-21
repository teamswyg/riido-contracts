package main

import (
	"path/filepath"
	"strings"
)

func requiredCommands(workflow, id string) []string {
	workflow = filepath.ToSlash(workflow)
	switch {
	case id == "ci" && strings.HasSuffix(workflow, ".github/workflows/ci.yml"):
		return []string{
			`test "$(go list -m all | wc -l | tr -d ' ')" = "1"`,
			"go run ./tools/enumgen verify",
			"go run ./tools/fsmgen verify",
			"go test ./...",
		}
	case id == "go-ci" && strings.HasSuffix(workflow, ".github/workflows/go-ci.yml"):
		return []string{
			"go mod download",
			"go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.12.2",
			"golangci-lint run ./... --timeout=5m",
			"go run ./tools/enumgen verify",
			"go run ./tools/fsmgen verify",
			"go test ./...",
		}
	default:
		return nil
	}
}
