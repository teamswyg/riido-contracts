package main

import "github.com/teamswyg/riido-contracts/ir"

func validRuntimeEvent() ir.CanonicalEvent {
	e := baseEvent(ir.EventScopeRuntime)
	e.Type = ir.EventRuntimeRegistered
	e.RuntimeID = "rt_1"
	return e
}

func validTaskEvent() ir.CanonicalEvent {
	e := baseEvent(ir.EventScopeTask)
	e.Type = ir.EventTaskQueued
	e.ActorKind = ir.ActorDaemon
	e.ActorID = "daemon-1"
	e.TaskID = "task_1"
	e.FSMVersion = 1
	return e
}
