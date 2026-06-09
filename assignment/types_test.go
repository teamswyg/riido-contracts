package assignment

import (
	"encoding/json"
	"testing"
	"time"
)

func TestAssignmentAPIJSONShapes(t *testing.T) {
	now := time.Date(2026, 5, 27, 11, 0, 0, 0, time.UTC)
	assignment := Assignment{
		ID:                       "asn-000001",
		TaskID:                   "task-a",
		ComponentID:              "component-1",
		AgentID:                  "agent-a",
		RuntimeProvider:          "codex",
		ModelID:                  "gpt-5.5",
		Prompt:                   "run tests",
		AgentInstruction:         "act as QA",
		AllowExperimentalRuntime: true,
		ResumeSessionID:          "th-prev",
		ProviderSessionID:        "th-current",
		State:                    AssignmentLeased,
		LeaseToken:               "lease-1",
		ReplacesAssignmentID:     "asn-old",
		BlockedByAssignmentID:    "asn-blocker",
		CreatedAt:                now,
		UpdatedAt:                now,
	}
	assertJSON(t, "assign request", AssignRequest{
		ComponentID:              "component-1",
		AgentID:                  "agent-a",
		RuntimeProvider:          "codex",
		ModelID:                  "gpt-5.5",
		Prompt:                   "run tests",
		AgentInstruction:         "act as QA",
		AllowExperimentalRuntime: true,
		ResumeSessionID:          "th-prev",
		CreatedBy:                "user-a",
	}, `{"component_id":"component-1","agent_id":"agent-a","runtime_provider":"codex","model_id":"gpt-5.5","prompt":"run tests","agent_instruction":"act as QA","allow_experimental_runtime":true,"resume_session_id":"th-prev","created_by":"user-a"}`)
	assertJSON(t, "poll request", PollRequest{
		DaemonID:  "daemon-a",
		DeviceID:  "device-a",
		RuntimeID: "daemon-a:agent:agent-a:codex",
	}, `{"daemon_id":"daemon-a","device_id":"device-a","runtime_id":"daemon-a:agent:agent-a:codex"}`)
	assertJSON(t, "poll response", PollResponse{
		SchemaVersion: SchemaVersion,
		Action:        PollStart,
		Assignment:    &assignment,
	}, `{"schema_version":"riido-ai-server.v1","action":"start","assignment":{"assignment_id":"asn-000001","task_id":"task-a","component_id":"component-1","agent_id":"agent-a","runtime_provider":"codex","model_id":"gpt-5.5","prompt":"run tests","agent_instruction":"act as QA","allow_experimental_runtime":true,"resume_session_id":"th-prev","provider_session_id":"th-current","state":"leased","lease_token":"lease-1","replaces_assignment_id":"asn-old","blocked_by_assignment_id":"asn-blocker","created_at":"2026-05-27T11:00:00Z","updated_at":"2026-05-27T11:00:00Z"}}`)
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
	}, `{"schema_version":"riido-ai-server.v1","refreshed_assignments":[{"assignment_id":"asn-000001","task_id":"task-a","component_id":"component-1","agent_id":"agent-a","runtime_provider":"codex","model_id":"gpt-5.5","prompt":"run tests","agent_instruction":"act as QA","allow_experimental_runtime":true,"resume_session_id":"th-prev","provider_session_id":"th-current","state":"leased","lease_token":"lease-1","replaces_assignment_id":"asn-old","blocked_by_assignment_id":"asn-blocker","created_at":"2026-05-27T11:00:00Z","updated_at":"2026-05-27T11:00:00Z"}]}`)
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
	event := TaskEvent{
		Seq:          1,
		TaskID:       "task-a",
		AssignmentID: "asn-000001",
		AgentID:      "agent-a",
		Type:         EventAssignmentRunning,
		State:        AssignmentRunning,
		Message:      "running",
		Metadata:     map[string]string{"step": "run"},
		At:           now,
	}
	assertJSON(t, "agent event response", AgentEventResponse{
		SchemaVersion: SchemaVersion,
		Assignment:    &assignment,
		Event:         event,
	}, `{"schema_version":"riido-ai-server.v1","assignment":{"assignment_id":"asn-000001","task_id":"task-a","component_id":"component-1","agent_id":"agent-a","runtime_provider":"codex","model_id":"gpt-5.5","prompt":"run tests","agent_instruction":"act as QA","allow_experimental_runtime":true,"resume_session_id":"th-prev","provider_session_id":"th-current","state":"leased","lease_token":"lease-1","replaces_assignment_id":"asn-old","blocked_by_assignment_id":"asn-blocker","created_at":"2026-05-27T11:00:00Z","updated_at":"2026-05-27T11:00:00Z"},"event":{"seq":1,"task_id":"task-a","assignment_id":"asn-000001","agent_id":"agent-a","type":"assignment_running","state":"running","message":"running","metadata":{"step":"run"},"at":"2026-05-27T11:00:00Z"}}`)
	assertJSON(t, "agent runtime binding", AgentRuntimeBinding{
		AgentID:         "agent-a",
		DaemonID:        "daemon-a",
		DeviceID:        "device-a",
		RuntimeID:       "daemon-a:agent:agent-a:codex",
		RuntimeProvider: "codex",
	}, `{"agent_id":"agent-a","daemon_id":"daemon-a","device_id":"device-a","runtime_id":"daemon-a:agent:agent-a:codex","runtime_provider":"codex"}`)
}

func assertJSON(t *testing.T, name string, value any, want string) {
	t.Helper()
	data, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal %s: %v", name, err)
	}
	if got := string(data); got != want {
		t.Fatalf("%s JSON = %s, want %s", name, got, want)
	}
}

// TestPollRequestWaitMsAdditive locks the backward-compatibility guarantee for
// the long-poll hint: an omitted WaitMs marshals byte-identical to the legacy
// PollRequest (so an old control plane sees no new field), and a set WaitMs
// round-trips on the wire.
func TestPollRequestWaitMsAdditive(t *testing.T) {
	legacy := `{"daemon_id":"daemon-a","device_id":"device-a","runtime_id":"rt-a"}`

	// Omitted -> byte-identical to the pre-wait_ms shape.
	assertJSON(t, "poll request without wait_ms", PollRequest{
		DaemonID:  "daemon-a",
		DeviceID:  "device-a",
		RuntimeID: "rt-a",
	}, legacy)

	// Set -> wait_ms present.
	assertJSON(t, "poll request with wait_ms", PollRequest{
		DaemonID:  "daemon-a",
		DeviceID:  "device-a",
		RuntimeID: "rt-a",
		WaitMs:    20000,
	}, `{"daemon_id":"daemon-a","device_id":"device-a","runtime_id":"rt-a","wait_ms":20000}`)

	// Round-trip: a legacy payload decodes with WaitMs == 0.
	var got PollRequest
	if err := json.Unmarshal([]byte(legacy), &got); err != nil {
		t.Fatalf("unmarshal legacy: %v", err)
	}
	if got.WaitMs != 0 {
		t.Fatalf("legacy WaitMs = %d, want 0", got.WaitMs)
	}
}
