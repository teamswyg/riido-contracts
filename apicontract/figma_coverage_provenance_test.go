package apicontract

import (
	"os"
	"strings"
	"testing"
)

func verifyFigmaCoverageProvenance(t *testing.T, stabilizedBy []string, docPath string) {
	t.Helper()
	want := []string{
		"teamswyg/riido-contracts#38",
		"teamswyg/riido-contracts#39",
		"teamswyg/riido-contracts#45",
		"teamswyg/riido-contracts#46",
		"teamswyg/riido-contracts#51",
		"teamswyg/riido-contracts#52",
		"teamswyg/riido-contracts#54",
		"teamswyg/riido-contracts#55",
		"teamswyg/riido-contracts#56",
		"teamswyg/riido-contracts#57",
		"teamswyg/riido-contracts#58",
		"teamswyg/riido-contracts#60",
		"teamswyg/riido-contracts#62",
		"teamswyg/riido-contracts#63",
		"teamswyg/riido-contracts#64",
		"teamswyg/riido-contracts#65",
		"teamswyg/riido-contracts#66",
		"teamswyg/riido-contracts#67",
	}
	if len(stabilizedBy) != len(want) {
		t.Fatalf("stabilized_by = %d entries, want %d: %+v", len(stabilizedBy), len(want), stabilizedBy)
	}
	for i := range want {
		if stabilizedBy[i] != want[i] {
			t.Fatalf("stabilized_by[%d] = %q, want %q; full list = %+v", i, stabilizedBy[i], want[i], stabilizedBy)
		}
	}
	doc, err := os.ReadFile(docPath)
	if err != nil {
		t.Fatalf("read coverage doc for provenance: %v", err)
	}
	docText := string(doc)
	for _, pr := range want {
		if !strings.Contains(docText, pr) {
			t.Fatalf("coverage doc must mention stabilization provenance %q", pr)
		}
	}
	if !strings.Contains(docText, "`stabilized_by`") ||
		!strings.Contains(docText, "downstream projection") {
		t.Fatalf("coverage doc must explain stabilized_by as the downstream projection mirror source")
	}
}
