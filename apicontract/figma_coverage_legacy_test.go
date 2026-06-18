package apicontract

import (
	"strings"
	"testing"
)

func verifySemanticLegacyWireframeCoverage(t *testing.T, entries []figmaCoverageEntry, inventory map[string]map[string]figmaCoverageNode, primaryEntries map[string]figmaCoverageEntry, docText string) {
	t.Helper()
	required := map[string]struct {
		name       string
		absorbedBy string
	}{
		"13:3789": {"런타임", "162:23090"},
		"86:9988": {"런타임", "162:23090"},
		"17:3551": {"에이전트", "432:37336"},
		"17:4231": {"에이전트 수정", "432:37336"},
		"84:9846": {"에이전트 추가", "432:37336"},
		"17:2871": {"데몬 상세", "162:23090"},
		"17:3111": {"런타임 상세", "162:23090"},
	}
	byNode := map[string]figmaCoverageEntry{}
	for _, entry := range entries {
		byNode[entry.NodeID] = entry
	}
	for nodeID, want := range required {
		entry, ok := byNode[nodeID]
		if !ok {
			t.Fatalf("semantic legacy Wireframe node %s %s must be promoted from inventory to non_ui_top_level_nodes", nodeID, want.name)
		}
		if entry.PageID != "0:1" {
			t.Fatalf("semantic legacy node %s page_id = %q, want 0:1", nodeID, entry.PageID)
		}
		if entry.Name != want.name {
			t.Fatalf("semantic legacy node %s name = %q, want %q", nodeID, entry.Name, want.name)
		}
		if entry.CoverageStatus != "covered" {
			t.Fatalf("semantic legacy node %s coverage_status = %q, want covered", nodeID, entry.CoverageStatus)
		}
		if entry.EvidenceKind != "figma_legacy_wireframe_section" {
			t.Fatalf("semantic legacy node %s evidence_kind = %q, want figma_legacy_wireframe_section", nodeID, entry.EvidenceKind)
		}
		if _, ok := inventory["0:1"][nodeID]; !ok {
			t.Fatalf("semantic legacy node %s is not present in loaded Wireframe inventory", nodeID)
		}
		if entry.AbsorbedByTopLevelNodeID != want.absorbedBy {
			t.Fatalf("semantic legacy node %s absorbed_by_top_level_node_id = %q, want %q", nodeID, entry.AbsorbedByTopLevelNodeID, want.absorbedBy)
		}
		absorbed, ok := primaryEntries[want.absorbedBy]
		if !ok {
			t.Fatalf("semantic legacy node %s absorbs into missing primary UI entry %s", nodeID, want.absorbedBy)
		}
		if absorbed.CoverageStatus != "covered" {
			t.Fatalf("semantic legacy node %s absorbs into non-covered primary entry %s", nodeID, want.absorbedBy)
		}
		if len(entry.GeneratedPaths) == 0 {
			t.Fatalf("semantic legacy node %s must name the generated paths inherited from its current UI entry", nodeID)
		}
		for _, generatedPath := range entry.GeneratedPaths {
			if !stringSliceContains(absorbed.GeneratedPaths, generatedPath) {
				t.Fatalf("semantic legacy node %s generated path %q is not covered by absorbed primary UI entry %s", nodeID, generatedPath, want.absorbedBy)
			}
		}
		facts := strings.Join(entry.CoveredFacts, "\n")
		if !strings.Contains(facts, "absorbed by the current UI") {
			t.Fatalf("semantic legacy node %s must explain current UI absorption: %q", nodeID, facts)
		}
		if !strings.Contains(docText, nodeID) || !strings.Contains(docText, want.absorbedBy) {
			t.Fatalf("coverage doc must mention semantic legacy node %s and absorbed primary entry %s", nodeID, want.absorbedBy)
		}
	}
}
