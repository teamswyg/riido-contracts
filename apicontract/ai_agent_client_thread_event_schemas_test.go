package apicontract

import "testing"

func verifyTaskThreadRecordSchema(t *testing.T, openAPI OpenAPISpec) {
	t.Helper()
	props := openAPISchemaProperties(t, openAPI, "AIAgentTaskThreadRecord")
	verifyResultAndFailureFields(t, props, "AIAgentTaskThreadRecord")
}

func verifyWorkStatusEventSchema(t *testing.T, openAPI OpenAPISpec) {
	t.Helper()
	props := openAPISchemaProperties(t, openAPI, "AgentWorkStatusChangedEvent")
	verifyResultAndFailureFields(t, props, "AgentWorkStatusChangedEvent")
}

func verifyResultAndFailureFields(t *testing.T, props map[string]any, schemaName string) {
	t.Helper()
	resultMessage, ok := props["result_message"].(map[string]any)
	if !ok || resultMessage["type"] != "string" {
		t.Fatalf("%s result_message schema = %#v", schemaName, props["result_message"])
	}
	failureDiagnostics, ok := props["failure_diagnostics"].(map[string]any)
	if !ok || failureDiagnostics["$ref"] != "#/components/schemas/AIAgentTaskThreadFailureDiagnostics" {
		t.Fatalf("%s failure_diagnostics schema = %#v", schemaName, props["failure_diagnostics"])
	}
}
