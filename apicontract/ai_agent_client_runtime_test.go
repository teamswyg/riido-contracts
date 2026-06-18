package apicontract

import "testing"

func (f aiAgentClientContractFixture) verifyRuntimeSchemas(t *testing.T) {
	t.Helper()
	verifyRuntimeRecordSchema(t, f.openAPI)
	verifyDeviceDaemonSchemas(t, f.openAPI)
	verifyRuntimeModelSchema(t, f.openAPI)
}

func verifyRuntimeRecordSchema(t *testing.T, openAPI OpenAPISpec) {
	t.Helper()
	runtimeRecord := openAPI.Components.Schemas["RuntimeRecord"]
	runtimeProps := openAPISchemaPropertiesFromSchema(t, runtimeRecord, "RuntimeRecord")
	runtimeRequired := openAPISchemaRequired(t, runtimeRecord, "RuntimeRecord")
	if !contains(runtimeRequired, "requires_experimental_opt_in") || !contains(runtimeRequired, "provider_version") {
		t.Fatalf("RuntimeRecord required = %#v", runtimeRecord["required"])
	}
	optIn, ok := runtimeProps["requires_experimental_opt_in"].(map[string]any)
	if !ok || optIn["type"] != "boolean" {
		t.Fatalf("RuntimeRecord requires_experimental_opt_in schema = %#v", runtimeProps["requires_experimental_opt_in"])
	}
	models, ok := runtimeProps["models"].(map[string]any)
	if !ok || models["type"] != "array" {
		t.Fatalf("RuntimeRecord models schema = %#v", runtimeProps["models"])
	}
	providerVersion, ok := runtimeProps["provider_version"].(map[string]any)
	if !ok || providerVersion["type"] != "string" {
		t.Fatalf("RuntimeRecord provider_version schema = %#v", runtimeProps["provider_version"])
	}
}
