package main

func assignedProfileMapShapePass(schemas []schema) bool {
	response := findSchema(schemas, "AssignedAgentProfileMapResponse")
	if response.Name == "" {
		return false
	}
	active, participants := false, false
	for _, prop := range response.Properties {
		if prop.Name == "assigned_agent_profiles" {
			active = prop.AdditionalPropertiesRef == "AssignedAgentProfile"
		}
		if prop.Name == "task_agent_participants" {
			participants = prop.AdditionalPropertiesRef == "TaskAgentParticipantCollection"
		}
	}
	return active && participants
}
