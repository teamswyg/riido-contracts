package apicontract

import "testing"

func verifyAgentClientRecordV2(t *testing.T, openAPI OpenAPISpec) {
	t.Helper()
	props := openAPISchemaProperties(t, openAPI, "AgentClientRecordV2")
	if _, ok := props["workspace_id"].(map[string]any); !ok {
		t.Fatalf("AgentClientRecordV2 workspace_id missing: %#v", props)
	}
	required := openAPISchemaRequired(t, openAPI.Components.Schemas["AgentClientRecordV2"], "AgentClientRecordV2")
	if !contains(required, "workspace_id") {
		t.Fatalf("AgentClientRecordV2 required = %#v", openAPI.Components.Schemas["AgentClientRecordV2"]["required"])
	}
}
