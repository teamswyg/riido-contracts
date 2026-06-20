package capability

import "testing"

func TestCatalogsReturnCopies(t *testing.T) {
	protocols := AllProtocolKinds()
	protocols[0] = "mutated"
	if AllProtocolKinds()[0] != ProtocolClaudeStreamJSON {
		t.Fatal("AllProtocolKinds returned mutable backing store")
	}
}

func TestCatalogCounts(t *testing.T) {
	if len(AllProtocolKinds()) != 6 {
		t.Fatalf("protocol count = %d, want 6", len(AllProtocolKinds()))
	}
	if len(AllCompatibilityStatuses()) != 4 {
		t.Fatalf("compatibility status count = %d, want 4", len(AllCompatibilityStatuses()))
	}
}
