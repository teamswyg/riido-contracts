package apicontract

import "testing"

func TestAgentCatalogDSLGeneratesIRAndOpenAPI(t *testing.T) {
	dsl := loadTestDSL(t, "fixtures/control-plane-agent-catalog.dsl.riido.json")
	ir, err := GenerateIR(dsl)
	if err != nil {
		t.Fatalf("GenerateIR: %v", err)
	}
	if ir.SchemaVersion != IRSchemaVersion || ir.SourceSchemaVersion != DSLSchemaVersion {
		t.Fatalf("IR versions = %q / %q", ir.SchemaVersion, ir.SourceSchemaVersion)
	}
	if got, want := len(ir.Operations), 5; got != want {
		t.Fatalf("IR operations = %d, want %d", got, want)
	}
	openAPI, err := GenerateOpenAPI(ir)
	if err != nil {
		t.Fatalf("GenerateOpenAPI: %v", err)
	}
	verifyAgentCatalogOpenAPI(t, openAPI)
}

func verifyAgentCatalogOpenAPI(t *testing.T, openAPI OpenAPISpec) {
	t.Helper()
	if openAPI.OpenAPI != OpenAPIVersion {
		t.Fatalf("OpenAPI version = %q", openAPI.OpenAPI)
	}
	list := openAPI.Paths["/v1/agent-catalog"]["get"]
	if list.OperationID != "listAgents" {
		t.Fatalf("list operation = %+v", list)
	}
	if len(list.RiidoScopes) != 1 || list.RiidoScopes[0] != "agent-catalog:read" {
		t.Fatalf("list scopes = %v", list.RiidoScopes)
	}
	update := openAPI.Paths["/v1/agent-catalog/{agent_id}"]["patch"]
	if len(update.Parameters) != 1 || update.Parameters[0].Name != "agent_id" {
		t.Fatalf("update path parameters = %+v", update.Parameters)
	}
	if update.RiidoRBAC != "agent_catalog_visibility.v1" {
		t.Fatalf("update rbac = %q", update.RiidoRBAC)
	}
}
