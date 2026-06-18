package apicontract

import (
	"strings"
	"testing"
)

func (f aiAgentClientContractFixture) verifyAssignableAgents(t *testing.T) {
	t.Helper()
	assignable := f.openAPI.Paths["/v1/client/ai-agent/tasks/{task_id}/assignable-agents"]["get"]
	if len(assignable.Parameters) != 1 || assignable.Parameters[0].Name != "task_id" {
		t.Fatalf("assignable-agent parameters = %#v", assignable.Parameters)
	}
	props := openAPISchemaProperties(t, f.openAPI, "AgentClientListResponse")
	assignableAgents, ok := props["agents"].(map[string]any)
	if !ok {
		t.Fatalf("AgentClientListResponse agents property missing: %#v", props["agents"])
	}
	description, ok := assignableAgents["description"].(string)
	if !ok || !strings.Contains(description, "agent_id") || !strings.Contains(description, "tasks.assign") {
		t.Fatalf("AgentClientListResponse agents description must explain task dropdown agent_id source: %q", description)
	}
}
