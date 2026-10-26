package apicontract

import "testing"

func TestAIAgentRunReadTitleIsOptionalAndNotInSSE(t *testing.T) {
	f := loadAIAgentClientContractFixture(t)
	schema := f.openAPI.Components.Schemas["AIAgentRunRead"]
	props := openAPISchemaPropertiesFromSchema(t, schema, "AIAgentRunRead")
	title, ok := props["title"].(map[string]any)
	if !ok || title["type"] != "string" {
		t.Fatal("GET read metadata needs a current, optional display title")
	}
	if contains(openAPISchemaRequired(t, schema, "AIAgentRunRead"), "title") {
		t.Fatal("sources without a current title must remain valid")
	}
	if _, exists := openAPISchemaProperties(t, f.openAPI, "AgentContextReadSource")["title"]; exists {
		t.Fatal("current GET titles must not become captured/replayed SSE fields")
	}
}
