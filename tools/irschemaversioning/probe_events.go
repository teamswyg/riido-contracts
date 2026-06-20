package main

import "github.com/teamswyg/riido-contracts/ir"

func runtimeWithoutRuntimeID() ir.CanonicalEvent {
	e := validRuntimeEvent()
	e.RuntimeID = ""
	return e
}

func taskWithRuntimeID() ir.CanonicalEvent {
	e := validTaskEvent()
	e.RuntimeID = "rt_1"
	return e
}

func runTransitionWithoutFSM() ir.CanonicalEvent {
	e := validRunEvent()
	e.Type = ir.EventTaskClaimed
	e.NativeConfigVersion = ""
	return e
}
