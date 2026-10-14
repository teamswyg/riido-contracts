package apicontract

import (
	"strings"
	"testing"
)

func verifyFigmaAPIGeneratedAnnotationContentPolicy(t *testing.T, policy figmaAPIGeneratedAnnotationContentRule, docText string) {
	t.Helper()
	if policy.CategoryID != "700:0" || policy.CategoryLabel != "API Generated" {
		t.Fatalf("API Generated annotation content category drifted: %+v", policy)
	}
	if len(policy.LabelFormat) != 3 {
		t.Fatalf("API Generated annotation label_format = %d entries, want 3", len(policy.LabelFormat))
	}
	verifyFigmaAPIGeneratedLabelFormat(t, policy, docText)
	if !strings.Contains(policy.Rule, "must not become a second API SSOT") {
		t.Fatalf("API Generated annotation content policy must prevent second SSOT drift: %q", policy.Rule)
	}
	verifyFigmaAPIGeneratedRetiredCategories(t, policy.RetiredCategories, docText)
	verifyFigmaAPIGeneratedLiveInspection(t, policy.LiveInspection, docText)
}

func verifyFigmaAPIGeneratedLabelFormat(t *testing.T, policy figmaAPIGeneratedAnnotationContentRule, docText string) {
	t.Helper()
	joined := strings.Join(policy.LabelFormat, "\n") + "\n" + policy.Rule
	for _, needle := range []string{"riido.*", "v2.", "source coverage entry", "종류", "Query", "Mutation", "SSE Stream", "배경", "Korean", "text/event-stream", "non-stream GET", "non-GET"} {
		if !strings.Contains(joined, needle) {
			t.Fatalf("API Generated annotation content policy must mention %q: %+v", needle, policy)
		}
		if !strings.Contains(docText, apiPolicyDocNeedle(needle)) {
			t.Fatalf("coverage doc must mention API Generated annotation content policy %q", needle)
		}
	}
}

func apiPolicyDocNeedle(needle string) string {
	switch needle {
	case "non-stream GET":
		return "non-stream `GET`"
	case "non-GET":
		return "non-`GET`"
	default:
		return needle
	}
}
