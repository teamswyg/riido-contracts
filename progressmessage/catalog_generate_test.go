package progressmessage

import "testing"

func TestProgressMessageDSLGeneratesIR(t *testing.T) {
	dsl := loadTestDSL(t)
	ir, err := GenerateIR(dsl)
	if err != nil {
		t.Fatalf("GenerateIR: %v", err)
	}
	assertFixture(t, "catalog.ir.riido.json", ir)
}
