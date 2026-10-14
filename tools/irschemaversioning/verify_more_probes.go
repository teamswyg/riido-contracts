package main

import (
	"errors"

	"github.com/teamswyg/riido-contracts/ir"
)

func verifyAdvancedProbes() error {
	if !hasViolation(ir.ValidateEnvelope(runTransitionWithoutFSM()), "INVALID_FSMVERSION", "FSMVersion") {
		return errors.New("RunScope transition without FSMVersion was not rejected")
	}
	if !hasViolation(ir.ValidateEnvelope(runWithFakePlaceholder()), "FAKE_PLACEHOLDER", "RuntimeID") {
		return errors.New("fake RuntimeID placeholder was not rejected")
	}
	if !hasViolation(phaseDependentAfterNCV(), "MISSING_FIELD", "NativeConfigVersion") {
		return errors.New("phase-dependent event after NCV establishment was not rejected")
	}
	return nil
}

func hasViolation(vs []ir.EnvelopeViolation, code, field string) bool {
	for _, v := range vs {
		if v.Code == code && v.Field == field {
			return true
		}
	}
	return false
}
