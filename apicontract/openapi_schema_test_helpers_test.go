package apicontract

import "testing"

func openAPISchemaProperties(t *testing.T, openAPI OpenAPISpec, schemaName string) map[string]any {
	t.Helper()
	return openAPISchemaPropertiesFromSchema(t, openAPI.Components.Schemas[schemaName], schemaName)
}

func openAPISchemaPropertiesFromSchema(t *testing.T, schema map[string]any, schemaName string) map[string]any {
	t.Helper()
	props, ok := schema["properties"].(map[string]any)
	if !ok {
		t.Fatalf("%s properties missing: %#v", schemaName, schema)
	}
	return props
}

func openAPISchemaRequired(t *testing.T, schema map[string]any, schemaName string) []string {
	t.Helper()
	required, ok := schema["required"].([]string)
	if !ok {
		t.Fatalf("%s required = %#v", schemaName, schema["required"])
	}
	return required
}
