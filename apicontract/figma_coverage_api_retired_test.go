package apicontract

import (
	"strings"
	"testing"
)

func verifyFigmaAPIGeneratedRetiredCategories(t *testing.T, categories []figmaAPIGeneratedAnnotationRetiredCategory, docText string) {
	t.Helper()
	if len(categories) != 1 {
		t.Fatalf("API Generated retired categories = %d, want 1", len(categories))
	}
	retired := categories[0]
	if retired.CategoryID != "39:0" || retired.CategoryLabel != "클라이언트 전달" {
		t.Fatalf("unexpected retired API Generated category: %+v", retired)
	}
	if retired.RetirementStatus != "unused_not_deleted" || retired.LiveUsageCount != 0 {
		t.Fatalf("retired API Generated category must stay unused_not_deleted with zero live usage: %+v", retired)
	}
	if retired.ObservedAt != "2026-06-03" || !strings.Contains(retired.ToolLimitation, "design owner") {
		t.Fatalf("retired API Generated category must record automation limitation: %+v", retired)
	}
	for _, needle := range []string{retired.CategoryID, retired.CategoryLabel, "retired", "zero annotations"} {
		if !strings.Contains(docText, needle) {
			t.Fatalf("coverage doc must mention retired API Generated category %q", needle)
		}
	}
}
