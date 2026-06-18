package apicontract

import (
	"strings"
	"testing"
)

func (s *figmaAPIGeneratedInventoryScope) verifyGroup(t *testing.T, group figmaAPIGeneratedAnnotationGroup) {
	t.Helper()
	if strings.TrimSpace(group.UIArea) == "" {
		t.Fatalf("API Generated annotation group has empty ui_area: %+v", group)
	}
	if group.CategoryID != "700:0" || group.CategoryLabel != "API Generated" {
		t.Fatalf("API Generated annotation group %q category drifted: %+v", group.FigmaGeneratedPath, group)
	}
	if !strings.HasPrefix(group.FigmaGeneratedPath, "riido.") {
		t.Fatalf("API Generated annotation group must preserve Figma facade path: %q", group.FigmaGeneratedPath)
	}
	canonical := canonicalPathFromFigmaFacade(group.FigmaGeneratedPath)
	if group.CanonicalGeneratedPath != canonical {
		t.Fatalf("API Generated annotation group %q canonical path = %q, want %q", group.FigmaGeneratedPath, group.CanonicalGeneratedPath, canonical)
	}
	s.verifyGroupPath(t, group)
	annotationCount := s.verifyGroupSources(t, group, "v2."+group.CanonicalGeneratedPath)
	if group.AnnotationCount != annotationCount {
		t.Fatalf("API Generated annotation group %q annotation_count = %d, want node count %d", group.CanonicalGeneratedPath, group.AnnotationCount, annotationCount)
	}
	s.totalAnnotations += annotationCount
	s.verifyGroupDocumentation(t, group)
}

func (s *figmaAPIGeneratedInventoryScope) verifyGroupDocumentation(t *testing.T, group figmaAPIGeneratedAnnotationGroup) {
	t.Helper()
	for _, needle := range []string{group.UIArea, group.FigmaGeneratedPath, group.CanonicalGeneratedPath, group.OperationKind, group.Background} {
		if !strings.Contains(s.docText, needle) {
			t.Fatalf("coverage doc must mention API Generated annotation inventory %q", needle)
		}
	}
	if !docMentionsGeneratedPath(s.docText, "v2."+group.CanonicalGeneratedPath) {
		t.Fatalf("coverage doc must mention API Generated annotation v2 counterpart %q", "v2."+group.CanonicalGeneratedPath)
	}
}
