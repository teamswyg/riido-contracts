package main

import "github.com/teamswyg/riido-contracts/ir"

func validRunEvent() ir.CanonicalEvent {
	e := baseEvent(ir.EventScopeRun)
	e.Type = ir.EventTextDelta
	e.ActorKind = ir.ActorAgent
	e.ActorID = "run_1"
	e.TaskID = "task_1"
	e.RunID = "run_1"
	e.RuntimeID = "rt_1"
	e.CapabilityFingerprint = "fp_abc"
	e.ProviderKind = "claude"
	e.ProtocolKind = "claude-stream-json"
	e.ProviderVersion = "2.1.128"
	e.AdapterID = "claude-stream-json"
	e.AdapterVersion = "0.1.0"
	e.ProtocolVersion = "stream-json-v1"
	e.NativeConfigVersion = "nc_xyz"
	return e
}
