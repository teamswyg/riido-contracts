package apicontract

import (
	"strings"
	"testing"
)

func verifyFigmaAPIGeneratedAnnotations(t *testing.T, annotations []figmaAPIGeneratedAnnotation, docText string, openAPIGeneratedPaths, registered map[string]string, entries map[string]figmaCoverageEntry) {
	t.Helper()
	if got, want := len(annotations), 2; got != want {
		t.Fatalf("api_generated_annotations = %d, want %d", got, want)
	}
	seen := map[string]bool{}
	for _, annotation := range annotations {
		if seen[annotation.NodeID] {
			t.Fatalf("duplicate API Generated annotation node %q", annotation.NodeID)
		}
		seen[annotation.NodeID] = true
		if _, ok := registered[annotation.NodeID]; !ok {
			t.Fatalf("API Generated annotation %q is not a registered Figma evidence node", annotation.NodeID)
		}
		if annotation.TopLevelNodeID != "153:15931" || annotation.CoverageEntryNodeID != "153:15931" {
			t.Fatalf("API Generated annotation %q must resolve through task-thread top-level entry 153:15931: %+v", annotation.NodeID, annotation)
		}
		if annotation.CategoryID != "700:0" || annotation.CategoryLabel != "API Generated" {
			t.Fatalf("API Generated annotation %q category drifted: %+v", annotation.NodeID, annotation)
		}
		if !strings.HasPrefix(annotation.FigmaGeneratedPath, "riido.") {
			t.Fatalf("API Generated annotation %q must preserve the Figma facade path: %q", annotation.NodeID, annotation.FigmaGeneratedPath)
		}
		canonical := canonicalPathFromFigmaFacade(annotation.FigmaGeneratedPath)
		if annotation.CanonicalGeneratedPath != canonical {
			t.Fatalf("API Generated annotation %q canonical path = %q, want %q", annotation.NodeID, annotation.CanonicalGeneratedPath, canonical)
		}
		if _, ok := openAPIGeneratedPaths[annotation.CanonicalGeneratedPath]; !ok {
			t.Fatalf("API Generated annotation %q references unknown OpenAPI generated path %q", annotation.NodeID, annotation.CanonicalGeneratedPath)
		}
		entry, ok := entries[annotation.CoverageEntryNodeID]
		if !ok {
			t.Fatalf("API Generated annotation %q references missing coverage entry %q", annotation.NodeID, annotation.CoverageEntryNodeID)
		}
		if !stringSliceContains(entry.GeneratedPaths, annotation.CanonicalGeneratedPath) {
			t.Fatalf("API Generated annotation %q canonical path %q is not in coverage entry %q generated paths", annotation.NodeID, annotation.CanonicalGeneratedPath, entry.NodeID)
		}
		for _, needle := range []string{annotation.NodeID, annotation.FigmaGeneratedPath, annotation.CanonicalGeneratedPath, annotation.CategoryLabel} {
			if !strings.Contains(docText, needle) {
				t.Fatalf("coverage doc must mention API Generated annotation %q", needle)
			}
		}
		if strings.Contains(annotation.FigmaLabel, "작업중") {
			if annotation.ResolutionStatus != "resolved_stale_handoff_copy" {
				t.Fatalf("API Generated annotation %q stale Figma copy must be explicitly resolved: %+v", annotation.NodeID, annotation)
			}
			if !strings.Contains(annotation.Resolution, "stale") || !strings.Contains(docText, "상세내용은 작업중입니다") {
				t.Fatalf("API Generated annotation %q stale copy resolution is not documented", annotation.NodeID)
			}
		}
	}
}
