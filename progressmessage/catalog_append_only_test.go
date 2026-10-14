package progressmessage

import "testing"

func TestProgressMessageCatalogIsAppendOnlyAgainstHEAD(t *testing.T) {
	previous, err := baselineCatalogFromGit()
	if err != nil {
		t.Skip("progress message catalog is new in this checkout")
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
