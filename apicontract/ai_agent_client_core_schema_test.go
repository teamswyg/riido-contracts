package apicontract

import "testing"

func (f aiAgentClientContractFixture) verifyCoreSchemas(t *testing.T) {
	t.Helper()
	runtimeAvailability := f.openAPI.Components.Schemas["RuntimeAvailability"]
	values, ok := runtimeAvailability["enum"].([]string)
	if !ok || len(values) != 2 || values[0] != "online" || values[1] != "offline" {
		t.Fatalf("RuntimeAvailability enum = %#v", runtimeAvailability["enum"])
	}
	streamEvent := f.openAPI.Components.Schemas["ClientStreamEvent"]
	if _, ok := streamEvent["oneOf"].([]map[string]any); !ok {
		t.Fatalf("ClientStreamEvent oneOf missing: %#v", streamEvent)
	}
	streamOperation := f.openAPI.Paths["/v1/client/ai-agent/events"]["get"]
	if _, ok := streamOperation.Responses["200"].Content["text/event-stream"]; !ok {
		t.Fatalf("stream response content = %#v", streamOperation.Responses["200"].Content)
	}
	if _, ok := f.openAPI.Components.Schemas["AgentOnboardingTemplate"]; ok {
		t.Fatalf("AgentOnboardingTemplate must not be exposed")
	}
}
