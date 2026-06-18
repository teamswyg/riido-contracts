package apicontract

import (
	"strings"
	"testing"
)

func (f aiAgentClientContractFixture) verifyBootstrapSchema(t *testing.T) {
	t.Helper()
	bootstrapSchema := f.openAPI.Components.Schemas["ClientBootstrapResponse"]
	bootstrapProps := openAPISchemaPropertiesFromSchema(t, bootstrapSchema, "ClientBootstrapResponse")
	if _, ok := bootstrapProps["agent_templates"]; ok {
		t.Fatalf("ClientBootstrapResponse must not expose agent_templates: %#v", bootstrapProps)
	}
	agents, ok := bootstrapProps["agents"].(map[string]any)
	if !ok {
		t.Fatalf("ClientBootstrapResponse agents property missing: %#v", bootstrapProps["agents"])
	}
	description, ok := agents["description"].(string)
	if !ok ||
		!strings.Contains(description, "agent_id") ||
		!strings.Contains(description, "tasks.assignableAgents") {
		t.Fatalf("ClientBootstrapResponse agents description = %q", description)
	}
}
