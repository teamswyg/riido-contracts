package apicontract

import (
	"strings"
	"testing"
)

func verifyHeadlessFileKeyLimitation(t *testing.T, limitation figmaSupportingToolLimitation, docText string) {
	t.Helper()
	if limitation.ID == "" {
		t.Fatalf("supporting_tool_limitations must include figma-headless-file-key-placeholder.v1")
	}
	for _, needle := range []string{"use_figma", "figma.fileKey"} {
		if !strings.Contains(limitation.Tool, needle) {
			t.Fatalf("headless file-key limitation tool must contain %q: %+v", needle, limitation)
		}
	}
	for _, needle := range []string{"MUOd9lctoEHASUStN3vUuK", "figma.fileKey=headless", "pages and annotation categories"} {
		if !strings.Contains(limitation.ObservedResult, needle) {
			t.Fatalf("headless file-key observed_result must contain %q: %q", needle, limitation.ObservedResult)
		}
	}
	for _, needle := range []string{"MUOd9lctoEHASUStN3vUuK", "v.1.22 AI Agent"} {
		if !stringSliceContains(limitation.AuthoritativeResult, needle) {
			t.Fatalf("headless file-key authoritative_result must contain %q: %+v", needle, limitation.AuthoritativeResult)
		}
	}
	rule := strings.ToLower(limitation.Rule)
	for _, needle := range []string{"supporting evidence only", "must not overwrite figma.file_key", "expected_pages", "downstream projection source identity"} {
		if !strings.Contains(rule, needle) {
			t.Fatalf("headless file-key rule must contain %q: %q", needle, limitation.Rule)
		}
	}
	for _, needle := range []string{"figma-headless-file-key-placeholder.v1", "`figma.fileKey=headless`", "`MUOd9lctoEHASUStN3vUuK`", "authoritative file identity", "must not overwrite `figma.file_key`"} {
		if !strings.Contains(docText, needle) {
			t.Fatalf("coverage doc must describe headless file-key limitation with %q", needle)
		}
	}
}
