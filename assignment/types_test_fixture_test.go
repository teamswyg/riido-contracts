package assignment

import (
	"encoding/json"
	"testing"
	"time"
)

func testAssignmentTime() time.Time {
	return time.Date(2026, 5, 27, 11, 0, 0, 0, time.UTC)
}

func testAssignmentWorktree() *AssignmentWorktree {
	return &AssignmentWorktree{
		RepositoryFullName: "teamswyg/riido-daemon",
		RepositoryURL:      "https://github.com/teamswyg/riido-daemon",
		BranchName:         "RIID-4964-agent-profile-upload",
		Source:             "connected_pull_request",
	}
}

func testAssignment() Assignment {
	now := testAssignmentTime()
	return Assignment{
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
		Worktree:                 testAssignmentWorktree(),
		State:                    AssignmentLeased,
		LeaseToken:               "lease-1",
		ReplacesAssignmentID:     "asn-old",
		BlockedByAssignmentID:    "asn-blocker",
		CreatedAt:                now,
		UpdatedAt:                now,
	}
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
