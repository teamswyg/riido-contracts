package main

import "fmt"

func verifyMapShape(model model) error {
	if !model.MapShapePass {
		return fmt.Errorf("assigned_agent_profiles must map to AssignedAgentProfile additional properties")
	}
	return nil
}
