package apicontract

import (
	"strings"
	"testing"
)

func verifyOnboardingPageLoadTimeoutLimitation(t *testing.T, limitation figmaSupportingToolLimitation, docText string) {
	t.Helper()
	if limitation.ID == "" {
		t.Fatalf("supporting_tool_limitations must include figma-onboarding-page-load-timeout.v1")
	}
	for _, needle := range []string{"get_metadata", "42:3014", "Plugin API page load"} {
		if !strings.Contains(limitation.Tool, needle) {
			t.Fatalf("onboarding load timeout tool must contain %q: %+v", needle, limitation)
		}
	}
	for _, needle := range []string{"time out after 120s", "Wireframe - 온보딩", "setCurrentPageAsync", "236:33845", "236:33847", "six onboarding riido.* API Generated annotations"} {
		if !strings.Contains(limitation.ObservedResult, needle) {
			t.Fatalf("onboarding load timeout observed_result must contain %q: %q", needle, limitation.ObservedResult)
		}
	}
	verifyOnboardingAuthoritativeResult(t, limitation)
	verifyOnboardingLimitationRule(t, limitation)
	for _, needle := range []string{"figma-onboarding-page-load-timeout.v1", "get_metadata(nodeId=42:3014)", "after 120s", "`Wireframe - 온보딩`", "`236:33845`", "`236:33847`", "six onboarding `riido.*` `API Generated`", "must not rewrite `expected_pages`", "onboarding generated paths unresolved", "`page.children.length=84`", "known captured inventory remains 83"} {
		if !strings.Contains(docText, needle) {
			t.Fatalf("coverage doc must describe onboarding page load timeout with %q", needle)
		}
	}
}

func verifyOnboardingAuthoritativeResult(t *testing.T, limitation figmaSupportingToolLimitation) {
	t.Helper()
	for _, needle := range []string{"42:3014", "child_count=84", "known_inventory_count=83", "unresolved_extra_top_level_node=1", "non_ui_top_level_inventory", "236:33845", "236:33847", "onboarding_api_generated_annotations=6"} {
		if !stringSliceContains(limitation.AuthoritativeResult, needle) {
			t.Fatalf("onboarding load timeout authoritative_result must contain %q: %+v", needle, limitation.AuthoritativeResult)
		}
	}
}

func verifyOnboardingLimitationRule(t *testing.T, limitation figmaSupportingToolLimitation) {
	t.Helper()
	rule := strings.ToLower(limitation.Rule)
	for _, needle := range []string{"supporting evidence only", "must not rewrite expected_pages", "remove page 42:3014", "onboarding generated paths", "direct registered-node lookup", "known_inventory_count may lag expected_pages.child_count"} {
		if !strings.Contains(rule, needle) {
			t.Fatalf("onboarding load timeout rule must contain %q: %q", needle, limitation.Rule)
		}
	}
}
