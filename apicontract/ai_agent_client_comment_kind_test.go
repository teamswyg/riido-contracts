package apicontract

import "testing"

func (f aiAgentClientContractFixture) verifyCommentKindSchema(t *testing.T) {
	t.Helper()
	commentKind := f.openAPI.Components.Schemas["AgentTaskCommentKind"]
	commentValues, ok := commentKind["enum"].([]string)
	if !ok || len(commentValues) == 0 || commentValues[0] != "queued_by_busy_agent" {
		t.Fatalf("AgentTaskCommentKind enum = %#v", commentKind["enum"])
	}
}
