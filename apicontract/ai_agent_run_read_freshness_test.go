package apicontract

import "testing"

func TestAIAgentRunReadFreshnessIsOptionalAndNotInSSE(t *testing.T) {
	f := loadAIAgentClientContractFixture(t)
	schema := f.openAPI.Components.Schemas["AIAgentRunRead"]
	props := openAPISchemaPropertiesFromSchema(t, schema, "AIAgentRunRead")
	field, ok := props["modified_since_read"].(map[string]any)
	if !ok || field["type"] != "boolean" {
		t.Fatal("GET reads need an optional current-source freshness comparison")
	}
	if contains(openAPISchemaRequired(t, schema, "AIAgentRunRead"), "modified_since_read") {
		t.Fatal("an unavailable comparison must be omitted, not reported as unchanged")
	}
	if _, exists := openAPISchemaProperties(t, f.openAPI, "AgentContextReadSource")["modified_since_read"]; exists {
		t.Fatal("current freshness comparisons must not be captured or replayed in SSE")
	}
}
