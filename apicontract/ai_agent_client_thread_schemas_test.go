package apicontract

import "testing"

func verifyTaskThreadSchemas(t *testing.T, openAPI OpenAPISpec) {
	t.Helper()
	threadCollectionProps := openAPISchemaProperties(t, openAPI, "AIAgentTaskThreadCollectionResponse")
	if _, ok := threadCollectionProps["active_stream"].(map[string]any); !ok {
		t.Fatalf("active_stream schema missing: %#v", threadCollectionProps["active_stream"])
	}
	verifyTaskActionResponseSchema(t, openAPI)
	verifyTaskThreadRecordSchema(t, openAPI)
	verifyWorkStatusEventSchema(t, openAPI)
	progressEvent := openAPI.Components.Schemas["AgentThreadProgressEvent"]
	progressRequired := openAPISchemaRequired(t, progressEvent, "AgentThreadProgressEvent")
	if !contains(progressRequired, "thread_id") {
		t.Fatalf("AgentThreadProgressEvent required = %#v", progressEvent["required"])
	}
}

func verifyTaskActionResponseSchema(t *testing.T, openAPI OpenAPISpec) {
	t.Helper()
	props := openAPISchemaProperties(t, openAPI, "AIAgentTaskActionResponse")
	verifyResultAndFailureFields(t, props, "AIAgentTaskActionResponse")
	activeStream, ok := props["active_stream"].(map[string]any)
	if !ok || activeStream["$ref"] != "#/components/schemas/AIAgentTaskThreadStreamLink" {
		t.Fatalf("AIAgentTaskActionResponse active_stream schema = %#v", props["active_stream"])
	}
}
