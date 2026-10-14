package assignment

import "testing"

func TestHeartbeatJSONShapes(t *testing.T) {
	assignment := testAssignment()
	assertJSON(t, "heartbeat request", AgentHeartbeatRequest{
		DaemonID:            "daemon-a",
		DeviceID:            "device-a",
		RuntimeID:           "daemon-a:agent:agent-a:codex",
		RunningTaskIDs:      []string{"task-a"},
		ActiveAssignmentIDs: []string{"asn-000001"},
	}, `{"daemon_id":"daemon-a","device_id":"device-a","runtime_id":"daemon-a:agent:agent-a:codex","running_task_ids":["task-a"],"active_assignment_ids":["asn-000001"]}`)
	assertJSON(t, "heartbeat response", AgentHeartbeatResponse{
		SchemaVersion:        SchemaVersion,
		RefreshedAssignments: []Assignment{assignment},
	}, `{"schema_version":"riido-ai-server.v1","refreshed_assignments":[{"assignment_id":"asn-000001","task_id":"task-a","component_id":"component-1","agent_id":"agent-a","runtime_provider":"codex","model_id":"gpt-5.5","prompt":"run tests","agent_instruction":"act as QA","allow_experimental_runtime":true,"resume_session_id":"th-prev","provider_session_id":"th-current","worktree":{"repository_full_name":"teamswyg/riido-daemon","repository_url":"https://github.com/teamswyg/riido-daemon","branch_name":"RIID-4964-agent-profile-upload","source":"connected_pull_request"},"state":"leased","lease_token":"lease-1","replaces_assignment_id":"asn-old","blocked_by_assignment_id":"asn-blocker","created_at":"2026-05-27T11:00:00Z","updated_at":"2026-05-27T11:00:00Z"}]}`)
}
