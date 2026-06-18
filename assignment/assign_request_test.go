package assignment

import "testing"

func TestAssignRequestJSONShape(t *testing.T) {
	assertJSON(t, "assign request", AssignRequest{
		ComponentID:              "component-1",
		AgentID:                  "agent-a",
		RuntimeProvider:          "codex",
		ModelID:                  "gpt-5.5",
		Prompt:                   "run tests",
		AgentInstruction:         "act as QA",
		AllowExperimentalRuntime: true,
		ResumeSessionID:          "th-prev",
		Worktree:                 testAssignmentWorktree(),
		CreatedBy:                "user-a",
	}, `{"component_id":"component-1","agent_id":"agent-a","runtime_provider":"codex","model_id":"gpt-5.5","prompt":"run tests","agent_instruction":"act as QA","allow_experimental_runtime":true,"resume_session_id":"th-prev","worktree":{"repository_full_name":"teamswyg/riido-daemon","repository_url":"https://github.com/teamswyg/riido-daemon","branch_name":"RIID-4964-agent-profile-upload","source":"connected_pull_request"},"created_by":"user-a"}`)
}
