package progressmessage

import (
	"encoding/json"
	"testing"
)

func TestProgressMessageCatalogIsAppendOnlyAgainstHEAD(t *testing.T) {
	out, err := baselineCatalogFromGit()
	if err != nil {
		t.Skip("progress message catalog is new in this checkout")
	}
	var previous IRDocument
	if err := json.Unmarshal(out, &previous); err != nil {
		t.Fatalf("decode previous catalog: %v", err)
	}
	current, err := Catalog()
	if err != nil {
		t.Fatalf("Catalog: %v", err)
	}
	assertCatalogAppendOnly(t, previous, current)
}

func assertCatalogAppendOnly(
	t *testing.T,
	previous IRDocument,
	current IRDocument,
) {
	t.Helper()
	currentByCode := map[int]MessageDefinition{}
	for _, message := range current.Messages {
		currentByCode[message.Code] = message
	}
	for _, old := range previous.Messages {
		assertMessageCodeAppendOnly(t, old, currentByCode[old.Code])
	}
}
