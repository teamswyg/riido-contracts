package apicontract

import "testing"

func verifyAgentConfigurationProperties(t *testing.T, schemaName string, props map[string]any) {
	t.Helper()
	for _, propertyName := range []string{"name", "profile_thumbnail_url", "description", "runtime_id", "model_id", "visibility", "instruction"} {
		if _, ok := props[propertyName].(map[string]any); !ok {
			t.Fatalf("%s missing %s: %#v", schemaName, propertyName, props)
		}
	}
	thumbnail, ok := props["profile_thumbnail_url"].(map[string]any)
	if !ok || thumbnail["format"] != "uri" {
		t.Fatalf("%s profile_thumbnail_url schema = %#v", schemaName, props["profile_thumbnail_url"])
	}
	description, ok := props["description"].(map[string]any)
	if !ok || description["maxLength"] != 160 {
		t.Fatalf("%s description schema = %#v", schemaName, props["description"])
	}
	instruction, ok := props["instruction"].(map[string]any)
	if !ok || instruction["maxLength"] != 1000 {
		t.Fatalf("%s instruction schema = %#v", schemaName, props["instruction"])
	}
}
