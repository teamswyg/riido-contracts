package apicontract

import "testing"

func verifyDeviceDaemonSchemas(t *testing.T, openAPI OpenAPISpec) {
	t.Helper()
	daemonRecord := openAPI.Components.Schemas["DeviceDaemonRecord"]
	daemonProps := openAPISchemaPropertiesFromSchema(t, daemonRecord, "DeviceDaemonRecord")
	daemonRequired := openAPISchemaRequired(t, daemonRecord, "DeviceDaemonRecord")
	if !contains(daemonRequired, "supported_actions") {
		t.Fatalf("DeviceDaemonRecord required = %#v", daemonRecord["required"])
	}
	appVersion, ok := daemonProps["app_version"].(map[string]any)
	if !ok || appVersion["type"] != "string" {
		t.Fatalf("DeviceDaemonRecord app_version schema = %#v", daemonProps["app_version"])
	}
	daemonStatusProps := openAPISchemaProperties(t, openAPI, "DeviceDaemonStatusEvent")
	if _, ok := daemonStatusProps["daemon"].(map[string]any); !ok {
		t.Fatalf("DeviceDaemonStatusEvent daemon schema missing: %#v", daemonStatusProps)
	}
}

func verifyRuntimeModelSchema(t *testing.T, openAPI OpenAPISpec) {
	t.Helper()
	runtimeModel := openAPI.Components.Schemas["RuntimeModelRecord"]
	modelRequired := openAPISchemaRequired(t, runtimeModel, "RuntimeModelRecord")
	if !contains(modelRequired, "model_id") || !contains(modelRequired, "is_default") {
		t.Fatalf("RuntimeModelRecord required = %#v", runtimeModel["required"])
	}
}
