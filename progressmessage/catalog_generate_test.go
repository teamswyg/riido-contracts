package progressmessage

import "testing"

func TestProgressMessageDSLGeneratesIR(t *testing.T) {
	dsl := loadTestDSL(t)
	ir, err := GenerateIR(dsl)
	if err != nil {
		t.Fatalf("GenerateIR: %v", err)
	}
	catalog, err := Catalog()
	if err != nil {
		t.Fatalf("Catalog: %v", err)
	}
	assertCatalogEqual(t, ir, catalog)
}
