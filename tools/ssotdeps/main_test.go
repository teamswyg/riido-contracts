package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDefaultManifestVerifies(t *testing.T) {
	var out bytes.Buffer
	if err := run([]string{"verify"}, &out); err != nil {
		t.Fatalf("verify default manifest: %v", err)
	}
	if got := out.String(); !strings.Contains(got, "verified 8 facts and 4 repo dependencies") {
		t.Fatalf("unexpected output: %q", got)
	}
}

func TestVerifyRejectsCycle(t *testing.T) {
	m := minimalManifest(t)
	cycle := repoDependency{
		ID:         "contracts-imports-control-plane",
		FromRepo:   "riido-contracts",
		ToRepo:     "riido-control-plane",
		FactIDs:    []string{"agent-concept"},
		LocalScope: "bad cycle",
	}
	m.RepoDependencies = append([]repoDependency{cycle}, m.RepoDependencies...)
	err := verifyManifest(m, testRoot(t))
	if err == nil || !strings.Contains(err.Error(), "cycle") {
		t.Fatalf("expected cycle error, got %v", err)
	}
}

func TestVerifyRejectsDuplicateFactID(t *testing.T) {
	m := minimalManifest(t)
	m.Facts = append(m.Facts, m.Facts[0])
	err := verifyManifest(m, testRoot(t))
	if err == nil || !strings.Contains(err.Error(), "facts must be sorted") && !strings.Contains(err.Error(), "duplicate fact id") {
		t.Fatalf("expected duplicate or sorting error, got %v", err)
	}
}

func TestVerifyRejectsMissingSourcePhrase(t *testing.T) {
	m := minimalManifest(t)
	m.Facts[0].SourceRefs[0].RequiredPhrase = "not present"
	err := verifyManifest(m, testRoot(t))
	if err == nil || !strings.Contains(err.Error(), "does not contain phrase") {
		t.Fatalf("expected missing phrase error, got %v", err)
	}
}

func minimalManifest(t *testing.T) manifest {
	t.Helper()
	return manifest{
		SchemaVersion: schemaVersion,
		ID:            "test-map",
		RiidoTask:     "RIID-TEST",
		HumanDoc:      "docs/30-architecture/ssot-dependency-map.md",
		Facts: []fact{
			{
				ID:             "agent-concept",
				Fact:           "Agent means a task-assignable abstraction of a configured runtime",
				HumanDocPhrase: "Agent means a task-assignable abstraction of a configured runtime",
				SourceRefs: []sourceRef{
					{
						Repo:           localRepo,
						Path:           "docs/20-domain/ai-agent-policy.md",
						RequiredPhrase: "Agent",
					},
				},
				Owner: ownerRef{
					Repo: localRepo,
					Path: "docs/20-domain/ai-agent-policy.md",
				},
				Downstreams: []downstream{
					{
						Repo:       "riido-control-plane",
						LocalScope: "test projection",
					},
				},
			},
		},
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

func testRoot(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	mustWrite(t, filepath.Join(root, "docs/30-architecture/ssot-dependency-map.md"), "Agent means a task-assignable abstraction of a configured runtime")
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
