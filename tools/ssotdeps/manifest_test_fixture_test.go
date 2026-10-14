package main

import (
	"os"
	"path/filepath"
	"testing"
)

func minimalManifest(t *testing.T) manifest {
	t.Helper()
	return manifest{
		SchemaVersion:    schemaVersion,
		ID:               "test-map",
		RiidoTask:        "RIID-TEST",
		Workflow:         ".github/workflows/architecture-docs.yml",
		EvidenceArtifact: "architecture-docs-evidence",
		Loop:             minimalLoop(),
		HumanDoc:         "docs/30-architecture/ssot-dependency-map.md",
		Facts:            []fact{minimalFact()},
		RepoDependencies: []repoDependency{
			{
				ID:         "control-plane-imports-contracts-policy",
				FromRepo:   "riido-control-plane",
				ToRepo:     localRepo,
				FactIDs:    []string{"agent-concept"},
				LocalScope: "test dependency",
			},
		},
	}
}

func minimalLoop() evidenceLoop {
	return evidenceLoop{
		Observation:   "observe",
		Hypothesis:    "hypothesis",
		Execute:       "execute",
		Evaluate:      "evaluate",
		Retrospective: "retro",
	}
}

func testRoot(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	docPath := filepath.Join(root, "docs/30-architecture/ssot-dependency-map.md")
	mustWrite(t, docPath, "Agent means a task-assignable abstraction of a configured runtime")
	mustWrite(t, filepath.Join(root, "docs/20-domain/ai-agent-policy.md"), "Agent")
	return root
}

func mustWrite(t *testing.T, path, body string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}
