package apicontract

import "testing"

func TestGenerateIRRejectsInvalidDSL(t *testing.T) {
	dsl := loadTestDSL(t, "fixtures/control-plane-agent-catalog.dsl.riido.json")
	dsl.SchemaVersion = "riido-api-dsl.v0"
	if _, err := GenerateIR(dsl); err == nil {
		t.Fatal("expected unsupported schema version error")
	}
}
