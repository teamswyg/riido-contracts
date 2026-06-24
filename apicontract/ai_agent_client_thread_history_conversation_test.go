package apicontract

import "testing"

func verifyHistoryRecordConversationFields(t *testing.T, openAPI OpenAPISpec) {
	t.Helper()
	record := openAPI.Components.Schemas["AIAgentTaskThreadHistoryRecord"]
	props := openAPISchemaPropertiesFromSchema(t, record, "AIAgentTaskThreadHistoryRecord")
	if props["conversation_id"].(map[string]any)["type"] != "string" ||
		props["parent_thread_id"].(map[string]any)["type"] != "string" {
		t.Fatalf("history conversation fields = %#v", props)
	}
	required := openAPISchemaRequired(t, record, "AIAgentTaskThreadHistoryRecord")
	if !contains(required, "conversation_id") {
		t.Fatalf("conversation_id must be required: %#v", required)
	}
}

func scenarioExists(ir IRDocument, scenarioID string) bool {
	for _, scenario := range ir.Scenarios {
		if scenario.ScenarioID == scenarioID {
			return true
		}
	}
	return false
}
