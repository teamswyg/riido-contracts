package main

import "testing"

func TestAssignedProfileEvidenceRequiresPersistentParticipantMap(t *testing.T) {
	for _, ref := range []string{"", "AssignedAgentProfile", "TaskAgentParticipantCollection"} {
		schemas := []schema{{Name: "AssignedAgentProfileMapResponse", Properties: []property{
			{Name: "assigned_agent_profiles", AdditionalPropertiesRef: "AssignedAgentProfile"},
			{Name: "task_agent_participants", AdditionalPropertiesRef: ref},
		}}}
		if got := assignedProfileMapShapePass(schemas); got != (ref == "TaskAgentParticipantCollection") {
			t.Fatalf("participant ref %q accepted=%v", ref, got)
		}
	}
}
