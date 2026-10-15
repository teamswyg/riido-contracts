package main

import "fmt"

func verifyMapShape(model model) error {
	if !model.MapShapePass {
		return fmt.Errorf("assigned_agent_profiles must map to AssignedAgentProfile and task_agent_participants to TaskAgentParticipantCollection")
	}
	return nil
}
