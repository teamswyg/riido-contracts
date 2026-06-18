package apicontract

import "testing"

func (f aiAgentClientContractFixture) verifyIdentity(t *testing.T) {
	t.Helper()
	if got, want := len(f.ir.Enums), 12; got != want {
		t.Fatalf("IR enums = %d, want %d", got, want)
	}
	if got, want := len(f.ir.SumTypes), 1; got != want {
		t.Fatalf("IR sum types = %d, want %d", got, want)
	}
	if f.ir.ContractID != "control-plane-ai-agent-client-api.v2" {
		t.Fatalf("IR contract_id = %q", f.ir.ContractID)
	}
	if f.openAPI.Info.Title != "control-plane-ai-agent-client-api.v2" {
		t.Fatalf("OpenAPI title = %q", f.openAPI.Info.Title)
	}
}

func (f aiAgentClientContractFixture) verifyClientModules(t *testing.T) {
	t.Helper()
	if len(f.openAPI.RiidoClientModules) != 2 ||
		f.openAPI.RiidoClientModules[0].Module != "aiAgent" ||
		f.openAPI.RiidoClientModules[1].Module != "v2" {
		t.Fatalf("client modules = %#v", f.openAPI.RiidoClientModules)
	}
	if _, ok := f.openAPI.Components.SecuritySchemes["riidoAIAgentToken"]; !ok {
		t.Fatalf("riidoAIAgentToken security scheme missing: %#v", f.openAPI.Components.SecuritySchemes)
	}
}
