package assignment

import "testing"

func TestAgentEventRequestJSONShape(t *testing.T) {
	assertJSON(t, "agent event request", AgentEventRequest{
		AssignmentID:      "asn-000001",
		TaskID:            "task-a",
		DaemonID:          "daemon-a",
		DeviceID:          "device-a",
		RuntimeID:         "daemon-a:agent:agent-a:codex",
		RuntimeProvider:   "codex",
		ModelID:           "gpt-5.5",
		ProviderSessionID: "th-current",
		State:             AssignmentRunning,
		EventType:         EventRiidoLog,
		Message:           "working",
		Metadata:          map[string]string{"step": "test"},
	}, `{"assignment_id":"asn-000001","task_id":"task-a","daemon_id":"daemon-a","device_id":"device-a","runtime_id":"daemon-a:agent:agent-a:codex","runtime_provider":"codex","model_id":"gpt-5.5","provider_session_id":"th-current","state":"running","event_type":"riido_log","message":"working","metadata":{"step":"test"}}`)
}

func TestAgentEventResponseJSONShape(t *testing.T) {
	assignment := testAssignment()
	assertJSON(t, "agent event response", AgentEventResponse{
		SchemaVersion: SchemaVersion,
		Assignment:    &assignment,
		Event:         testTaskEvent(),
	}, `{"schema_version":"riido-ai-server.v1","assignment":{"assignment_id":"asn-000001","task_id":"task-a","component_id":"component-1","agent_id":"agent-a","runtime_provider":"codex","model_id":"gpt-5.5","prompt":"run tests","agent_instruction":"act as QA","allow_experimental_runtime":true,"resume_session_id":"th-prev","provider_session_id":"th-current","worktree":{"repository_full_name":"teamswyg/riido-daemon","repository_url":"https://github.com/teamswyg/riido-daemon","branch_name":"RIID-4964-agent-profile-upload","source":"connected_pull_request"},"state":"leased","lease_token":"lease-1","replaces_assignment_id":"asn-old","blocked_by_assignment_id":"asn-blocker","created_at":"2026-05-27T11:00:00Z","updated_at":"2026-05-27T11:00:00Z"},"event":{"seq":1,"task_id":"task-a","assignment_id":"asn-000001","agent_id":"agent-a","type":"assignment_running","state":"running","message":"running","metadata":{"step":"run"},"at":"2026-05-27T11:00:00Z"}}`)
}
