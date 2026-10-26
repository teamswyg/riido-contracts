package apicontract

import "testing"

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
