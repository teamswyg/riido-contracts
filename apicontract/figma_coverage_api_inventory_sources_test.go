package apicontract

import (
	"strings"
	"testing"
)

func (s *figmaAPIGeneratedInventoryScope) verifyGroupSources(t *testing.T, group figmaAPIGeneratedAnnotationGroup, v2Path string) int {
	t.Helper()
	if len(group.Sources) == 0 {
		t.Fatalf("API Generated annotation group %q must name sources", group.CanonicalGeneratedPath)
	}
	annotationCount := 0
	for _, source := range group.Sources {
		s.verifyGroupSource(t, group, source, v2Path)
		annotationCount += s.registerSourceNodeIDs(t, group, source)
	}
	return annotationCount
}

func (s *figmaAPIGeneratedInventoryScope) verifyGroupSource(t *testing.T, group figmaAPIGeneratedAnnotationGroup, source figmaAPIGeneratedAnnotationSource, v2Path string) {
	t.Helper()
	if strings.TrimSpace(source.PageID) == "" || strings.TrimSpace(source.TopLevelNodeID) == "" || strings.TrimSpace(source.CoverageEntryNodeID) == "" {
		t.Fatalf("API Generated annotation group %q has invalid source: %+v", group.CanonicalGeneratedPath, source)
	}
	entry, ok := s.entries[source.CoverageEntryNodeID]
	if !ok {
		t.Fatalf("API Generated annotation group %q references missing coverage entry %q", group.CanonicalGeneratedPath, source.CoverageEntryNodeID)
	}
	if !stringSliceContains(entry.GeneratedPaths, group.CanonicalGeneratedPath) {
		t.Fatalf("API Generated annotation group %q canonical path is not covered by source entry %q", group.CanonicalGeneratedPath, source.CoverageEntryNodeID)
	}
	if !stringSliceContains(entry.GeneratedPaths, v2Path) {
		t.Fatalf("API Generated annotation group %q v2 counterpart %q is not covered by source entry %q", group.CanonicalGeneratedPath, v2Path, source.CoverageEntryNodeID)
	}
	registerFigmaNodeIDIfAbsent(t, s.registered, source.TopLevelNodeID, "api_generated_annotation_inventory top-level "+group.CanonicalGeneratedPath)
}

func (s *figmaAPIGeneratedInventoryScope) registerSourceNodeIDs(t *testing.T, group figmaAPIGeneratedAnnotationGroup, source figmaAPIGeneratedAnnotationSource) int {
	t.Helper()
	if len(source.NodeIDs) == 0 {
		t.Fatalf("API Generated annotation group %q source %q must list node_ids", group.CanonicalGeneratedPath, source.TopLevelNodeID)
	}
	sourceSeen := map[string]bool{}
	for _, nodeID := range source.NodeIDs {
		if strings.TrimSpace(nodeID) == "" {
			t.Fatalf("API Generated annotation group %q source has empty node id", group.CanonicalGeneratedPath)
		}
		if sourceSeen[nodeID] {
			t.Fatalf("API Generated annotation group %q source %q duplicates node %q", group.CanonicalGeneratedPath, source.TopLevelNodeID, nodeID)
		}
		sourceSeen[nodeID] = true
		registerFigmaNodeIDIfAbsent(t, s.registered, nodeID, "api_generated_annotation_inventory "+group.CanonicalGeneratedPath)
	}
	return len(source.NodeIDs)
}
