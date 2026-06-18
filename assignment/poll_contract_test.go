package assignment

import (
	"encoding/json"
	"testing"
)

func TestPollRequestJSONShape(t *testing.T) {
	assertJSON(t, "poll request", PollRequest{
		DaemonID:  "daemon-a",
		DeviceID:  "device-a",
		RuntimeID: "daemon-a:agent:agent-a:codex",
	}, `{"daemon_id":"daemon-a","device_id":"device-a","runtime_id":"daemon-a:agent:agent-a:codex"}`)
}

func TestPollResponseJSONShape(t *testing.T) {
	assignment := testAssignment()
	assertJSON(t, "poll response", PollResponse{
		SchemaVersion: SchemaVersion,
		Action:        PollStart,
		Assignment:    &assignment,
	}, `{"schema_version":"riido-ai-server.v1","action":"start","assignment":{"assignment_id":"asn-000001","task_id":"task-a","component_id":"component-1","agent_id":"agent-a","runtime_provider":"codex","model_id":"gpt-5.5","prompt":"run tests","agent_instruction":"act as QA","allow_experimental_runtime":true,"resume_session_id":"th-prev","provider_session_id":"th-current","worktree":{"repository_full_name":"teamswyg/riido-daemon","repository_url":"https://github.com/teamswyg/riido-daemon","branch_name":"RIID-4964-agent-profile-upload","source":"connected_pull_request"},"state":"leased","lease_token":"lease-1","replaces_assignment_id":"asn-old","blocked_by_assignment_id":"asn-blocker","created_at":"2026-05-27T11:00:00Z","updated_at":"2026-05-27T11:00:00Z"}}`)
}

func TestPollRequestWaitMsAdditive(t *testing.T) {
	legacy := `{"daemon_id":"daemon-a","device_id":"device-a","runtime_id":"rt-a"}`
	assertJSON(t, "poll request without wait_ms", PollRequest{
		DaemonID:  "daemon-a",
		DeviceID:  "device-a",
		RuntimeID: "rt-a",
	}, legacy)
	assertJSON(t, "poll request with wait_ms", PollRequest{
		DaemonID:  "daemon-a",
		DeviceID:  "device-a",
		RuntimeID: "rt-a",
		WaitMs:    20000,
	}, `{"daemon_id":"daemon-a","device_id":"device-a","runtime_id":"rt-a","wait_ms":20000}`)
	var got PollRequest
	if err := json.Unmarshal([]byte(legacy), &got); err != nil {
		t.Fatalf("unmarshal legacy: %v", err)
	}
	if got.WaitMs != 0 {
		t.Fatalf("legacy WaitMs = %d, want 0", got.WaitMs)
	}
}
