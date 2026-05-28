package apicontract

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"
)

func TestAgentCatalogDSLGeneratesIRAndOpenAPI(t *testing.T) {
	dsl := loadTestDSL(t)
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

func TestAgentCatalogGeneratedFixturesDoNotDrift(t *testing.T) {
	dsl := loadTestDSL(t)
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

func TestGenerateIRRejectsInvalidDSL(t *testing.T) {
	dsl := loadTestDSL(t)
	dsl.SchemaVersion = "riido-api-dsl.v0"
	if _, err := GenerateIR(dsl); err == nil {
		t.Fatal("expected unsupported schema version error")
	}
}

func loadTestDSL(t *testing.T) DSLDocument {
	t.Helper()
	data, err := os.ReadFile("fixtures/control-plane-agent-catalog.dsl.riido.json")
	if err != nil {
		t.Fatalf("read DSL fixture: %v", err)
	}
	var dsl DSLDocument
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&dsl); err != nil {
		t.Fatalf("decode DSL fixture: %v", err)
	}
	return dsl
}

func assertFixture(t *testing.T, path string, value any) {
	t.Helper()
	want, err := MarshalCanonical(value)
	if err != nil {
		t.Fatalf("marshal fixture: %v", err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("%s drifted; run go run ./tools/apicontract generate", path)
	}
}
