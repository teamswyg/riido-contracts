package apicontract

import (
	"strings"
	"testing"
)

func verifyFigmaRuntimeEndpointLabel(t *testing.T, evidence []figmaCoverageNode, runtimeEntry figmaCoverageEntry, docText string) {
	t.Helper()
	var found bool
	for _, node := range evidence {
		if node.NodeID == "129:17930" {
			found = true
			if !strings.Contains(strings.ToLower(node.Name), "endpoint") {
				t.Fatalf("runtime endpoint-looking evidence node must explain its role: %+v", node)
			}
		}
	}
	if !found {
		t.Fatal("runtime settings endpoint-looking label node-id=129:17930 must be registered as verified evidence")
	}
	if runtimeEntry.NodeID != "162:23090" {
		t.Fatalf("runtime settings coverage entry missing: %+v", runtimeEntry)
	}
	facts := strings.Join(runtimeEntry.CoveredFacts, "\n")
	normalizedDocText := strings.Join(strings.Fields(docText), " ")
	for _, needle := range []string{
		"node-id=129:17930",
		"not a canonical base URL",
		"generated path",
		"live host export",
	} {
		if !strings.Contains(facts, needle) {
			t.Fatalf("runtime settings coverage must classify endpoint-looking Figma label with %q: %q", needle, facts)
		}
		if !strings.Contains(normalizedDocText, needle) {
			t.Fatalf("coverage doc must mention runtime endpoint-looking label rule with %q", needle)
		}
	}
}
