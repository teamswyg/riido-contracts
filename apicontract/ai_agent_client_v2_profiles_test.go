package apicontract

import "testing"

func (f aiAgentClientContractFixture) verifyAssignedProfiles(t *testing.T) {
	t.Helper()
	assignedProfiles := f.openAPI.Paths["/v2/client/workspaces/{workspace_id}/ai-agent/tasks/assigned-agent-profiles"]["get"]
	if assignedProfiles.OperationID != "listWorkspaceAssignedAgentProfilesV2" ||
		assignedProfiles.RiidoClient == nil ||
		assignedProfiles.RiidoClient.GeneratedPath != "v2.aiAgent.tasks.assignedAgentProfiles" ||
		assignedProfiles.RiidoClient.CacheTag != "v2.aiAgent.tasks.assignedAgentProfiles" ||
		assignedProfiles.RiidoRBAC != "workspace_assigned_agent_profile_map.v1" {
		t.Fatalf("assigned profiles operation = %#v", assignedProfiles)
	}
	if len(assignedProfiles.Parameters) != 1 || assignedProfiles.Parameters[0].Name != "workspace_id" {
		t.Fatalf("assigned profiles parameters = %#v", assignedProfiles.Parameters)
	}
	props := openAPISchemaProperties(t, f.openAPI, "AssignedAgentProfileMapResponse")
	assignedProfileValues, ok := props["assigned_agent_profiles"].(map[string]any)
	if !ok {
		t.Fatalf("assigned_agent_profiles property missing: %#v", props)
	}
	additional, ok := assignedProfileValues["additionalProperties"].(map[string]any)
	if !ok || additional["$ref"] != "#/components/schemas/AssignedAgentProfile" {
		t.Fatalf("assigned_agent_profiles additionalProperties = %#v", assignedProfileValues["additionalProperties"])
	}
}
