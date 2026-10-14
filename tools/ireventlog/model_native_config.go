package main

import "github.com/teamswyg/riido-contracts/ir"

func countNativeConfig(events []ir.EventType) nativeConfigCounts {
	var counts nativeConfigCounts
	for _, event := range events {
		switch ir.NativeConfigRequirementOf(event) {
		case ir.NativeConfigForbidden:
			counts.Forbidden++
		case ir.NativeConfigOptionalPreExecute:
			counts.PreExecute++
		case ir.NativeConfigRequired:
			counts.Required++
		case ir.NativeConfigPhaseDependent:
			counts.PhaseDependent++
		}
	}
	return counts
}
