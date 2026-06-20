package main

import "fmt"

func verifyFSM(model model) error {
	m := model.Manifest
	checks := map[string][2]int{
		"fsm schema version": {model.FSMSchema, m.ExpectedFSMSchemaVersion},
		"state count":        {len(model.States), m.ExpectedStateCount},
		"start state count":  {len(model.StartStates), m.ExpectedStartStateCount},
		"terminal count":     {len(model.TerminalStates), m.ExpectedTerminalStateCount},
		"transition count":   {model.TransitionCount, m.ExpectedTransitionCount},
	}
	for name, values := range checks {
		if values[0] != values[1] {
			return fmt.Errorf("%s = %d, want %d", name, values[0], values[1])
		}
	}
	return nil
}
