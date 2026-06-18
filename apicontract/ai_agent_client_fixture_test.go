package apicontract

import "testing"

type aiAgentClientContractFixture struct {
	ir      IRDocument
	openAPI OpenAPISpec
}

func loadAIAgentClientContractFixture(t *testing.T) aiAgentClientContractFixture {
	t.Helper()
	dsl := loadTestDSL(t, "fixtures/control-plane-ai-agent-client.dsl.riido.json")
	ir, err := GenerateIR(dsl)
	if err != nil {
		t.Fatalf("GenerateIR: %v", err)
	}
	openAPI, err := GenerateOpenAPI(ir)
	if err != nil {
		t.Fatalf("GenerateOpenAPI: %v", err)
	}
	return aiAgentClientContractFixture{ir: ir, openAPI: openAPI}
}
