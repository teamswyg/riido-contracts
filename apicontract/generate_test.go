package apicontract

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"
)

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

func TestGenerateIRRejectsInvalidDSL(t *testing.T) {
	dsl := loadTestDSL(t, "fixtures/control-plane-agent-catalog.dsl.riido.json")
	dsl.SchemaVersion = "riido-api-dsl.v0"
	if _, err := GenerateIR(dsl); err == nil {
		t.Fatal("expected unsupported schema version error")
	}
}

func TestAIAgentClientDSLKeepsEnumsAndSumTypesCodegenSafe(t *testing.T) {
	dsl := loadTestDSL(t, "fixtures/control-plane-ai-agent-client.dsl.riido.json")
	ir, err := GenerateIR(dsl)
	if err != nil {
		t.Fatalf("GenerateIR: %v", err)
	}
	if got, want := len(ir.Enums), 9; got != want {
		t.Fatalf("IR enums = %d, want %d", got, want)
	}
	if got, want := len(ir.SumTypes), 1; got != want {
		t.Fatalf("IR sum types = %d, want %d", got, want)
	}
	openAPI, err := GenerateOpenAPI(ir)
	if err != nil {
		t.Fatalf("GenerateOpenAPI: %v", err)
	}
	runtimeAvailability := openAPI.Components.Schemas["RuntimeAvailability"]
	values, ok := runtimeAvailability["enum"].([]string)
	if !ok || len(values) != 2 || values[0] != "online" || values[1] != "offline" {
		t.Fatalf("RuntimeAvailability enum = %#v", runtimeAvailability["enum"])
	}
	streamEvent := openAPI.Components.Schemas["ClientStreamEvent"]
	if _, ok := streamEvent["oneOf"].([]map[string]any); !ok {
		t.Fatalf("ClientStreamEvent oneOf missing: %#v", streamEvent)
	}
	streamOperation := openAPI.Paths["/v1/client/ai-agent/events"]["get"]
	if _, ok := streamOperation.Responses["200"].Content["text/event-stream"]; !ok {
		t.Fatalf("stream response content = %#v", streamOperation.Responses["200"].Content)
	}
	editability := openAPI.Paths["/v1/client/ai-agent/agents/{agent_id}/editability"]["get"]
	if editability.RiidoRBAC != "agent_mutation_safety.v1" {
		t.Fatalf("editability rbac = %q", editability.RiidoRBAC)
	}
	assignable := openAPI.Paths["/v1/client/ai-agent/tasks/{task_id}/assignable-agents"]["get"]
	if len(assignable.Parameters) != 1 || assignable.Parameters[0].Name != "task_id" {
		t.Fatalf("assignable-agent parameters = %#v", assignable.Parameters)
	}
	threads := openAPI.Paths["/v1/client/ai-agent/tasks/{task_id}/threads"]["get"]
	if threads.RiidoRBAC != "task_thread_cold_collection.v1" {
		t.Fatalf("task threads rbac = %q", threads.RiidoRBAC)
	}
	if len(threads.Parameters) != 1 || threads.Parameters[0].Name != "task_id" {
		t.Fatalf("task threads parameters = %#v", threads.Parameters)
	}
	commentKind := openAPI.Components.Schemas["AgentTaskCommentKind"]
	commentValues, ok := commentKind["enum"].([]string)
	if !ok || len(commentValues) == 0 || commentValues[0] != "queued_by_busy_agent" {
		t.Fatalf("AgentTaskCommentKind enum = %#v", commentKind["enum"])
	}
	agentRecord := openAPI.Components.Schemas["AgentClientRecord"]
	recordProps, ok := agentRecord["properties"].(map[string]any)
	if !ok {
		t.Fatalf("AgentClientRecord properties missing: %#v", agentRecord)
	}
	thumbnail, ok := recordProps["profile_thumbnail_url"].(map[string]any)
	if !ok || thumbnail["format"] != "uri" {
		t.Fatalf("profile_thumbnail_url schema = %#v", recordProps["profile_thumbnail_url"])
	}
	description, ok := recordProps["description"].(map[string]any)
	if !ok || description["maxLength"] != 160 {
		t.Fatalf("description schema = %#v", recordProps["description"])
	}
	instruction, ok := recordProps["instruction"].(map[string]any)
	if !ok || instruction["maxLength"] != 1000 {
		t.Fatalf("instruction schema = %#v", recordProps["instruction"])
	}
	if _, ok := recordProps["model_id"].(map[string]any); !ok {
		t.Fatalf("model_id schema missing: %#v", recordProps)
	}
	runtimeRecord := openAPI.Components.Schemas["RuntimeRecord"]
	runtimeProps, ok := runtimeRecord["properties"].(map[string]any)
	if !ok {
		t.Fatalf("RuntimeRecord properties missing: %#v", runtimeRecord)
	}
	models, ok := runtimeProps["models"].(map[string]any)
	if !ok || models["type"] != "array" {
		t.Fatalf("RuntimeRecord models schema = %#v", runtimeProps["models"])
	}
	runtimeModel := openAPI.Components.Schemas["RuntimeModelRecord"]
	modelRequired, ok := runtimeModel["required"].([]string)
	if !ok || !contains(modelRequired, "model_id") || !contains(modelRequired, "is_default") {
		t.Fatalf("RuntimeModelRecord required = %#v", runtimeModel["required"])
	}
	threadCollection := openAPI.Components.Schemas["AIAgentTaskThreadCollectionResponse"]
	threadCollectionProps, ok := threadCollection["properties"].(map[string]any)
	if !ok {
		t.Fatalf("AIAgentTaskThreadCollectionResponse properties missing: %#v", threadCollection)
	}
	if _, ok := threadCollectionProps["active_stream"].(map[string]any); !ok {
		t.Fatalf("active_stream schema missing: %#v", threadCollectionProps["active_stream"])
	}
	progressEvent := openAPI.Components.Schemas["AgentThreadProgressEvent"]
	progressRequired, ok := progressEvent["required"].([]string)
	if !ok || !contains(progressRequired, "thread_id") {
		t.Fatalf("AgentThreadProgressEvent required = %#v", progressEvent["required"])
	}
}

func TestAIAgentClientGeneratedFixturesDoNotDrift(t *testing.T) {
	dsl := loadTestDSL(t, "fixtures/control-plane-ai-agent-client.dsl.riido.json")
	ir, err := GenerateIR(dsl)
	if err != nil {
		t.Fatalf("GenerateIR: %v", err)
	}
	openAPI, err := GenerateOpenAPI(ir)
	if err != nil {
		t.Fatalf("GenerateOpenAPI: %v", err)
	}
	assertFixture(t, "fixtures/control-plane-ai-agent-client.ir.riido.json", ir)
	assertFixture(t, "fixtures/control-plane-ai-agent-client.openapi.json", openAPI)
}

func loadTestDSL(t *testing.T, path string) DSLDocument {
	t.Helper()
	data, err := os.ReadFile(path)
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

func contains(values []string, needle string) bool {
	for _, value := range values {
		if value == needle {
			return true
		}
	}
	return false
}
