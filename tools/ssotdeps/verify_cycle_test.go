package main

import (
	"strings"
	"testing"
)

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
