package assignment

import (
	"encoding/json"
	"testing"
)

func TestExecutionIdentityFromAssignment(t *testing.T) {
	assignment := Assignment{
		ID:                " asn-1 ",
		TaskID:            " task-a ",
		ComponentID:       " component-a ",
		AgentID:           " agent-a ",
		ResumeSessionID:   " sess-old ",
		ProviderSessionID: " sess-current ",
	}
	identity := IdentityFromAssignment(assignment, " runtime-a ")
	if identity.ExecutionID() != "asn-1" ||
		identity.RunID != "asn-1" ||
		identity.TaskID != "task-a" ||
		identity.ComponentID != "component-a" ||
		identity.AgentID != "agent-a" ||
		identity.RuntimeID != "runtime-a" ||
		identity.ProviderSessionID != "sess-current" ||
		!identity.Valid() {
		t.Fatalf("identity = %+v", identity)
	}
	assertJSON(t, "execution identity", identity, `{"assignment_id":"asn-1","task_id":"task-a","component_id":"component-a","agent_id":"agent-a","runtime_id":"runtime-a","run_id":"asn-1","provider_session_id":"sess-current"}`)
}

func TestExecutionIdentityFallbacks(t *testing.T) {
	assignment := Assignment{TaskID: " task-a ", AgentID: " agent-a ", ResumeSessionID: " sess-old "}
	identity := IdentityFromAssignment(assignment, "")
	if identity.ExecutionID() != "task-a" ||
		identity.RunID != "task-a" ||
		identity.ProviderSessionID != "sess-old" ||
		!identity.Valid() {
		t.Fatalf("fallback identity = %+v", identity)
	}
	identity.TaskID = ""
	if identity.Valid() {
		t.Fatalf("identity without task_id must be invalid: %+v", identity)
	}
}

func TestExecutionIdentityJSONShape(t *testing.T) {
	data, err := json.Marshal(ExecutionIdentity{
		AssignmentID: "asn-1",
		TaskID:       "task-a",
		AgentID:      "agent-a",
		RunID:        "run-a",
	})
	if err != nil {
		t.Fatal(err)
	}
	if got, want := string(data), `{"assignment_id":"asn-1","task_id":"task-a","agent_id":"agent-a","run_id":"run-a"}`; got != want {
		t.Fatalf("execution identity JSON = %s, want %s", got, want)
	}
}
