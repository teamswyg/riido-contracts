package apicontract

import "testing"

func TestAIAgentClientDisplayProgressIsOptionalReadOnlyCopy(t *testing.T) {
	fixture := loadAIAgentClientContractFixture(t)
	for _, name := range []string{"AgentThreadProgressLine", "AIAgentTaskThreadHistoryMessage"} {
		t.Run(name, func(t *testing.T) {
			properties := openAPISchemaProperties(t, fixture.openAPI, name)
			field, ok := properties["display_message"].(map[string]any)
			if !ok || field["type"] != "string" {
				t.Fatalf("%s must expose optional display_message string: %#v", name, field)
			}
			schema := fixture.openAPI.Components.Schemas[name]
			for _, property := range openAPISchemaRequired(t, schema, name) {
				if property == "display_message" {
					t.Fatal("display_message must remain optional for older servers")
				}
			}
			for _, internal := range []string{"message_code", "message_args"} {
				if _, exists := properties[internal]; exists {
					t.Fatalf("%s must not expose internal %s", name, internal)
				}
			}
		})
	}
}
