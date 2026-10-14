package apicontract

import (
	"strings"
	"testing"
)

func verifyAgentRecordIDDescription(t *testing.T, props map[string]any) {
	t.Helper()
	agentID, ok := props["agent_id"].(map[string]any)
	if !ok {
		t.Fatalf("agent_id schema missing: %#v", props)
	}
	description, ok := agentID["description"].(string)
	if !ok || !strings.Contains(description, "bootstrap.agents[]") || !strings.Contains(description, "tasks.assignableAgents.agents[]") {
		t.Fatalf("agent_id description must explain bootstrap vs assignableAgents source: %q", description)
	}
}

func verifyAgentRecordTimestamps(t *testing.T, props map[string]any) {
	t.Helper()
	createdAt, ok := props["created_at"].(map[string]any)
	if !ok || createdAt["format"] != "date-time" {
		t.Fatalf("created_at schema = %#v", props["created_at"])
	}
	updatedAt, ok := props["updated_at"].(map[string]any)
	if !ok || updatedAt["format"] != "date-time" {
		t.Fatalf("updated_at schema = %#v", props["updated_at"])
	}
}
