package apicontract

import (
	"strings"
	"testing"
)

func verifyFigmaCoverageInspectionMethod(t *testing.T, method figmaCoverageInspectionMethod, docText string) {
	t.Helper()
	if method.ID != "figma-plugin-api-page-registry.v1" {
		t.Fatalf("inspection_method.id = %q", method.ID)
	}
	if strings.TrimSpace(method.Authority) != "Figma Plugin API via use_figma" {
		t.Fatalf("inspection_method.authority = %q", method.Authority)
	}
	if method.PageRegistryExpression != "figma.root.children" {
		t.Fatalf("inspection_method.page_registry_expression = %q", method.PageRegistryExpression)
	}
	if method.TopLevelChildCountExpression != "await figma.setCurrentPageAsync(page); page.children.length" {
		t.Fatalf("inspection_method.top_level_child_count_expression = %q", method.TopLevelChildCountExpression)
	}
	if len(method.SupportingTools) == 0 {
		t.Fatalf("inspection_method.supporting_tools must name non-authoritative read tools")
	}
	rule := strings.ToLower(method.Rule)
	for _, needle := range []string{"metadata", "supporting evidence", "must not redefine page-level child counts", "lazy/unloaded"} {
		if !strings.Contains(rule, needle) {
			t.Fatalf("inspection_method.rule must contain %q: %q", needle, method.Rule)
		}
	}
	for _, needle := range []string{"figma.root.children", "await figma.setCurrentPageAsync(page)", "page.children.length", "Metadata XML/read", "supporting evidence only", "lazy/unloaded"} {
		if !strings.Contains(docText, needle) {
			t.Fatalf("coverage doc must describe inspection method with %q", needle)
		}
	}
}
