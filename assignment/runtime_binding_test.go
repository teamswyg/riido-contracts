package assignment

import "testing"

func TestAgentRuntimeBindingJSONShape(t *testing.T) {
	assertJSON(t, "agent runtime binding", AgentRuntimeBinding{
		AgentID:         "agent-a",
		DaemonID:        "daemon-a",
		DeviceID:        "device-a",
		RuntimeID:       "daemon-a:agent:agent-a:codex",
		RuntimeProvider: "codex",
	}, `{"agent_id":"agent-a","daemon_id":"daemon-a","device_id":"device-a","runtime_id":"daemon-a:agent:agent-a:codex","runtime_provider":"codex"}`)
}
