package apicontract

import (
	"strings"
	"testing"
)

func verifyMetadataPageListLimitation(t *testing.T, limitation figmaSupportingToolLimitation, docText string) {
	t.Helper()
	if limitation.ID == "" {
		t.Fatalf("supporting_tool_limitations must include figma-metadata-page-list-underreports-pages.v1")
	}
	if !strings.Contains(limitation.Tool, "get_metadata") || !strings.Contains(limitation.Tool, "without nodeId") {
		t.Fatalf("metadata limitation tool must name no-nodeId get_metadata: %+v", limitation)
	}
	for _, needle := range []string{"only page 129:5215 UI", "MUOd9lctoEHASUStN3vUuK"} {
		if !strings.Contains(limitation.ObservedResult, needle) {
			t.Fatalf("metadata limitation observed_result must contain %q: %q", needle, limitation.ObservedResult)
		}
	}
	verifyMetadataPageListAuthoritativePages(t, limitation)
	rule := strings.ToLower(limitation.Rule)
	for _, needle := range []string{"supporting evidence only", "must not remove expected_pages", "non-ui inventories"} {
		if !strings.Contains(rule, needle) {
			t.Fatalf("metadata limitation rule must contain %q: %q", needle, limitation.Rule)
		}
	}
	for _, needle := range []string{"figma-metadata-page-list-underreports-pages.v1", "get_metadata", "without `nodeId`", "`129:5215`", "`42:3014`", "`0:1`", "must not remove `expected_pages`"} {
		if !strings.Contains(docText, needle) {
			t.Fatalf("coverage doc must describe metadata page-list limitation with %q", needle)
		}
	}
}

func verifyMetadataPageListAuthoritativePages(t *testing.T, limitation figmaSupportingToolLimitation) {
	t.Helper()
	requiredPages := map[string]bool{"129:5215": false, "42:3014": false, "0:1": false}
	for _, pageID := range limitation.AuthoritativeResult {
		if _, ok := requiredPages[pageID]; ok {
			requiredPages[pageID] = true
		}
	}
	for pageID, seen := range requiredPages {
		if !seen {
			t.Fatalf("metadata limitation authoritative_result is missing page %s: %+v", pageID, limitation.AuthoritativeResult)
		}
	}
}
