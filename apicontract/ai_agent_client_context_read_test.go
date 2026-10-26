package apicontract

import (
	"slices"
	"testing"
)

func TestContextReadStreamContract(t *testing.T) {
	fixture := loadAIAgentClientContractFixture(t)
	props := openAPISchemaProperties(t, fixture.openAPI, "AgentContextReadEvent")
	for _, name := range []string{"event_type", "agent_id", "assignment_id", "thread_id", "run_id", "via", "served_at", "sources"} {
		if _, ok := props[name]; !ok {
			t.Errorf("AgentContextReadEvent missing %s", name)
		}
	}
	for _, name := range []string{"requester_id", "request_id", "body", "url"} {
		if _, ok := props[name]; ok {
			t.Errorf("private field %s leaked into stream", name)
		}
	}
	sources := openAPISchemaProperties(t, fixture.openAPI, "AgentContextReadSource")
	for _, name := range []string{"kind", "id", "part", "version", "updated_at", "status"} {
		if _, ok := sources[name]; !ok {
			t.Errorf("AgentContextReadSource missing %s", name)
		}
	}
	for _, name := range []string{"body", "url", "requester_id", "request_id"} {
		if _, ok := sources[name]; ok {
			t.Errorf("private field %s leaked into source", name)
		}
	}
}

func TestAIAgentClientAgentDirectoryReadsMatchHTTPAndSSESchema(t *testing.T) {
	fixture := loadAIAgentClientContractFixture(t)
	// Agent directory reads already emitted by control-plane HTTP history and SSE.
	for _, name := range []string{"AIAgentRunRead", "AgentContextReadSource"} {
		props := openAPISchemaProperties(t, fixture.openAPI, name)
		for field, value := range map[string]string{"kind": "agent", "part": "profile", "status": "served"} {
			property, ok := props[field].(map[string]any)
			if !ok {
				t.Fatalf("%s.%s missing", name, field)
			}
			values, ok := property["enum"].([]string)
			if !ok || !slices.Contains(values, value) {
				t.Errorf("%s.%s rejects emitted value %q", name, field, value)
			}
		}
	}
}
