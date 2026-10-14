package main

import "github.com/teamswyg/riido-contracts/task"

func stateRows(fsm task.TaskFSM) []stateRow {
	states := fsm.States()
	out := make([]stateRow, len(states))
	for i, state := range states {
		out[i] = stateRow{Name: state.String(), Kind: pointKind(fsm.PointKind(state))}
	}
	return out
}

func pointKind(kind task.TaskFSMPointKind) string {
	switch kind {
	case task.TaskFSMPointStart:
		return "start"
	case task.TaskFSMPointEnd:
		return "terminal"
	case task.TaskFSMPointIntermediate:
		return "intermediate"
	default:
		return "unknown"
	}
}
