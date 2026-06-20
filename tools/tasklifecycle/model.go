package main

import "github.com/teamswyg/riido-contracts/task"

func buildModel(m manifest) model {
	fsm := task.GeneratedTaskFSM()
	return model{
		Manifest:        m,
		FSMSchema:       task.FSMSchemaVersion,
		States:          stateRows(fsm),
		StartStates:     stateNames(fsm.StartStates()),
		TerminalStates:  stateNames(fsm.TerminalStates()),
		TransitionCount: len(fsm.Transitions()),
		Transitions:     groupTransitions(fsm.Transitions()),
	}
}

func stateNames(states []task.TaskStateCode) []string {
	out := make([]string, len(states))
	for i, state := range states {
		out[i] = state.String()
	}
	return out
}
