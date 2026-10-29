package apicontract

import "testing"

func TestAIAgentRunReadGroupingIsResponseOnly(t *testing.T) {
	f := loadAIAgentClientContractFixture(t)
	read := f.openAPI.Components.Schemas["AIAgentRunRead"]
	props := openAPISchemaPropertiesFromSchema(t, read, "AIAgentRunRead")
	for _, name := range []string{"owner_kind", "owner_id"} {
		field, ok := props[name].(map[string]any)
		if !ok || field["type"] != "string" {
			t.Fatalf("GET reads need optional %s for grouping child sources", name)
		}
		if contains(openAPISchemaRequired(t, read, "AIAgentRunRead"), name) {
			t.Fatalf("%s must be omitted when a read is its own group", name)
		}
	}
	listed, ok := props["listed"].(map[string]any)
	if !ok || listed["type"] != "boolean" || !contains(openAPISchemaRequired(t, read, "AIAgentRunRead"), "listed") {
		t.Fatal("every read must state whether the read UI lists it")
	}
	response := f.openAPI.Components.Schemas["AIAgentRunReadsResponse"]
	total, ok := openAPISchemaPropertiesFromSchema(t, response, "AIAgentRunReadsResponse")["source_total"].(map[string]any)
	if !ok || total["type"] != "integer" || !contains(openAPISchemaRequired(t, response, "AIAgentRunReadsResponse"), "source_total") {
		t.Fatal("GET reads need a required viewer-visible source_total across pages")
	}
	sse := openAPISchemaProperties(t, f.openAPI, "AgentContextReadSource")
	for _, name := range []string{"owner_kind", "owner_id", "listed", "source_total"} {
		if _, exists := sse[name]; exists {
			t.Fatalf("%s is computed for the current viewer and must not be replayed in SSE", name)
		}
	}
}
