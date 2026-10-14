package main

func assignedProfileMapShapePass(schemas []schema) bool {
	response := findSchema(schemas, "AssignedAgentProfileMapResponse")
	if response.Name == "" {
		return false
	}
	for _, prop := range response.Properties {
		if prop.Name == "assigned_agent_profiles" {
			return prop.AdditionalPropertiesRef == "AssignedAgentProfile"
		}
	}
	return false
}
