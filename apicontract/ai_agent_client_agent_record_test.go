package apicontract

import "testing"

func (f aiAgentClientContractFixture) verifyAgentRecordSchemas(t *testing.T) {
	t.Helper()
	f.verifyDaemonContracts(t)
	props := openAPISchemaProperties(t, f.openAPI, "AgentClientRecord")
	verifyAgentRecordProfileFields(t, props)
	verifyAgentRecordIDDescription(t, props)
	verifyAgentRecordTimestamps(t, props)
}

func verifyAgentRecordProfileFields(t *testing.T, props map[string]any) {
	t.Helper()
	thumbnail, ok := props["profile_thumbnail_url"].(map[string]any)
	if !ok || thumbnail["format"] != "uri" {
		t.Fatalf("profile_thumbnail_url schema = %#v", props["profile_thumbnail_url"])
	}
	description, ok := props["description"].(map[string]any)
	if !ok || description["maxLength"] != 160 {
		t.Fatalf("description schema = %#v", props["description"])
	}
	instruction, ok := props["instruction"].(map[string]any)
	if !ok || instruction["maxLength"] != 1000 {
		t.Fatalf("instruction schema = %#v", props["instruction"])
	}
	if _, ok := props["model_id"].(map[string]any); !ok {
		t.Fatalf("model_id schema missing: %#v", props)
	}
	if _, ok := props["tmp_color"].(map[string]any); !ok {
		t.Fatalf("tmp_color schema missing: %#v", props)
	}
}
