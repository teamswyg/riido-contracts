package main

import "fmt"

func verifyEntry(entry coverageEntry) error {
	if blank(entry.NodeID) || blank(entry.Name) || blank(entry.CoverageStatus) || blank(entry.EvidenceKind) {
		return fmt.Errorf("node_id, name, coverage_status, and evidence_kind are required")
	}
	switch entry.CoverageStatus {
	case "covered", "no_diff_product_surface":
		return verifyCoveredEntry(entry)
	case "non_decision_asset":
		if blank(entry.Reason) {
			return fmt.Errorf("%s entry requires reason", entry.CoverageStatus)
		}
		return nil
	default:
		return fmt.Errorf("unsupported coverage_status %q", entry.CoverageStatus)
	}
}

func verifyCoveredEntry(entry coverageEntry) error {
	if len(entry.SSOTDocs) == 0 || len(entry.OwnerRepos) == 0 || len(entry.CoveredFacts) == 0 {
		return fmt.Errorf("covered entry requires ssot_docs, owner_repos, and covered_facts")
	}
	if blank(entry.DirectionLoop.TopDown) || blank(entry.DirectionLoop.BottomUp) {
		return fmt.Errorf("covered entry requires top_down and bottom_up loop evidence")
	}
	return nil
}
