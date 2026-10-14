package apicontract

import (
	"strings"
	"testing"
)

func assertCoverageDocMentionsEntry(t *testing.T, docText string, entry figmaCoverageEntry) {
	t.Helper()
	if !strings.Contains(docText, entry.NodeID) || !strings.Contains(docText, entry.Name) {
		t.Fatalf("coverage doc must mention node %s %s", entry.NodeID, entry.Name)
	}
	for _, generatedPath := range entry.GeneratedPaths {
		if !docMentionsGeneratedPath(docText, generatedPath) {
			t.Fatalf("coverage doc must mention generated path %q for node %s", generatedPath, entry.NodeID)
		}
	}
}
