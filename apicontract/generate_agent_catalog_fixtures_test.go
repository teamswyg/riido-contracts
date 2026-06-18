package apicontract

import "testing"

func TestAgentCatalogGeneratedFixturesDoNotDrift(t *testing.T) {
	dsl := loadTestDSL(t, "fixtures/control-plane-agent-catalog.dsl.riido.json")
	ir, err := GenerateIR(dsl)
	if err != nil {
		t.Fatalf("GenerateIR: %v", err)
	}
	openAPI, err := GenerateOpenAPI(ir)
	if err != nil {
		t.Fatalf("GenerateOpenAPI: %v", err)
	}
	assertFixture(t, "fixtures/control-plane-agent-catalog.ir.riido.json", ir)
	assertFixture(t, "fixtures/control-plane-agent-catalog.openapi.json", openAPI)
}
