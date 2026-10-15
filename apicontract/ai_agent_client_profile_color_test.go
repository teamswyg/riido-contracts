package apicontract

import "testing"

func TestAIAgentClientProfileColorContract(t *testing.T) {
	f := loadAIAgentClientContractFixture(t)
	verifyAgentProfileColorEnum(t, f.openAPI)
	verifyAgentProfileColorProjections(t, f.openAPI)
}

func verifyAgentProfileColorEnum(t *testing.T, openAPI OpenAPISpec) {
	t.Helper()
	schema, ok := openAPI.Components.Schemas["AgentProfileColor"]
	if !ok {
		t.Fatalf("AgentProfileColor schema missing")
	}
	values, ok := schema["enum"].([]string)
	if !ok {
		t.Fatalf("AgentProfileColor enum = %#v", schema["enum"])
	}
	want := []string{"blue", "orange", "green", "purple"}
	if len(values) != len(want) {
		t.Fatalf("AgentProfileColor enum = %#v, want %#v", values, want)
	}
	for i := range want {
		if values[i] != want[i] {
			t.Fatalf("AgentProfileColor enum = %#v, want %#v", values, want)
		}
	}
}

func verifyAgentProfileColorProjections(t *testing.T, openAPI OpenAPISpec) {
	t.Helper()
	for _, name := range []string{"AgentClientRecord", "AgentClientRecordV2", "AssignedAgentProfile", "TaskAgentParticipantRecord", "AIAgentTaskThreadAgentSnapshot"} {
		props := openAPISchemaProperties(t, openAPI, name)
		if _, ok := props["profile_color"].(map[string]any); !ok {
			t.Fatalf("%s profile_color property missing: %#v", name, props)
		}
	}
	for _, name := range []string{"AgentClientRecordV2", "TaskAgentParticipantRecord"} {
		required := openAPISchemaRequired(t, openAPI.Components.Schemas[name], name)
		if !contains(required, "profile_color") {
			t.Fatalf("%s required must include profile_color: %#v", name, required)
		}
	}
	// Create/update request bodies must not accept a client-chosen color.
	for _, name := range []string{"CreateAgentConfigurationRequest", "UpdateAgentConfigurationRequest"} {
		if _, ok := openAPISchemaProperties(t, openAPI, name)["profile_color"]; ok {
			t.Fatalf("%s must not expose profile_color", name)
		}
	}
}
