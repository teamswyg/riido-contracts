package apicontract

import "testing"

func TestAIAgentParticipantRemovalGuardIsAvailableOnGeneratedWire(t *testing.T) {
	f := loadAIAgentClientContractFixture(t)
	path := "/v2/client/workspaces/{workspace_id}/ai-agent/tasks/{task_id}/agent-assignments/{agent_id}"
	removal := f.openAPI.Paths[path]["delete"]
	if removal.RequestBody == nil {
		t.Fatal("generated participant DELETE cannot carry the observed assignment guard")
	}
	if removal.RequestBody.Required {
		t.Fatal("participant DELETE must still accept requests without a body")
	}
	if ref := removal.RequestBody.Content["application/json"].Schema["$ref"]; ref != "#/components/schemas/DeleteAIAgentTaskAgentAssignmentRequest" {
		t.Fatalf("participant DELETE request ref = %s", ref)
	}
	for _, name := range []string{"DeleteAIAgentTaskAgentAssignmentRequest", "UnassignAIAgentTaskRequest"} {
		props := openAPISchemaProperties(t, f.openAPI, name)
		if _, ok := props["assignment_id"]; !ok {
			t.Fatalf("%s omits the observed removal guard", name)
		}
		if _, ok := props["reason"]; !ok {
			t.Fatalf("%s omits the removal reason", name)
		}
		required, _ := f.openAPI.Components.Schemas[name]["required"].([]string)
		for _, field := range required {
			if field == "assignment_id" || field == "reason" {
				t.Fatalf("%s requires optional removal metadata %s", name, field)
			}
		}
		if name == "DeleteAIAgentTaskAgentAssignmentRequest" {
			if _, ok := props["agent_id"]; ok {
				t.Fatal("path-Agent removal must not add an agent_id body input")
			}
		}
	}
}

func TestAIAgentParticipantRemovalRaceBelongsToMutationScenarios(t *testing.T) {
	f := loadAIAgentClientContractFixture(t)
	const scenario = ".interleaved-execution-prevents-participant-removal"
	for _, id := range []string{"unassignAIAgentTask", "unassignAIAgentTaskV2", "deleteAIAgentTaskAgentAssignmentV2"} {
		if !scenarioExists(f.ir, id+scenario) {
			t.Errorf("%s does not own its participant-removal race scenario", id)
		}
	}
	if scenarioExists(f.ir, "listWorkspaceAssignedAgentProfilesV2"+scenario) {
		t.Error("the GET read model still owns the DELETE commit race")
	}
}
