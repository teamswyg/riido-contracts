package apicontract

import (
	"strings"
	"testing"
)

func verifyCoverageEntry(t *testing.T, entry figmaCoverageEntry, openAPIGeneratedPaths map[string]string) {
	t.Helper()
	if strings.TrimSpace(entry.CoverageStatus) == "" {
		t.Fatalf("entry %q coverage_status is required", entry.NodeID)
	}
	switch entry.CoverageStatus {
	case "covered", "no_diff_product_surface", "planning_evidence":
		if len(entry.SSOTDocs) == 0 {
			t.Fatalf("entry %q must link ssot_docs", entry.NodeID)
		}
		if len(entry.OwnerRepos) == 0 {
			t.Fatalf("entry %q must name owner_repos", entry.NodeID)
		}
		if strings.TrimSpace(entry.DirectionLoop.TopDown) == "" || strings.TrimSpace(entry.DirectionLoop.BottomUp) == "" {
			t.Fatalf("entry %q must define both direction loops", entry.NodeID)
		}
		for _, doc := range entry.SSOTDocs {
			assertCoverageLocalRefExists(t, doc)
		}
		for _, generatedPath := range entry.GeneratedPaths {
			if _, ok := openAPIGeneratedPaths[generatedPath]; !ok {
				t.Fatalf("entry %q references unknown generated path %q", entry.NodeID, generatedPath)
			}
		}
	case "non_decision_asset":
		if strings.TrimSpace(entry.Reason) == "" {
			t.Fatalf("non-decision entry %q must explain reason", entry.NodeID)
		}
		if len(entry.SSOTDocs) != 0 || len(entry.OwnerRepos) != 0 {
			t.Fatalf("non-decision entry %q must not invent owners or SSOT docs", entry.NodeID)
		}
	default:
		t.Fatalf("entry %q has unknown coverage_status %q", entry.NodeID, entry.CoverageStatus)
	}
}
