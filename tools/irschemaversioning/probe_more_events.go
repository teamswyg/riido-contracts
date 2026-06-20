package main

import "github.com/teamswyg/riido-contracts/ir"

func runWithFakePlaceholder() ir.CanonicalEvent {
	e := validRunEvent()
	e.RuntimeID = "unknown"
	return e
}

func phaseDependentAfterNCV() []ir.EnvelopeViolation {
	e := validRunEvent()
	e.Type = ir.EventTaskFailed
	e.FSMVersion = 1
	e.NativeConfigVersion = ""
	return ir.ValidateEnvelopeWithRunContext(e, ir.RunContext{NativeConfigEstablished: true})
}
