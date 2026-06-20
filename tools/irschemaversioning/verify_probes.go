package main

import (
	"errors"

	"github.com/teamswyg/riido-contracts/ir"
)

func verifyValidatorProbes() error {
	for _, event := range []ir.CanonicalEvent{validSystemEvent(), validRuntimeEvent(), validTaskEvent(), validRunEvent()} {
		if violations := ir.ValidateEnvelope(event); len(violations) != 0 {
			return errors.New("valid scope fixture failed envelope validation")
		}
	}
	if err := verifyNegativeProbes(); err != nil {
		return err
	}
	return verifyAdvancedProbes()
}

func verifyNegativeProbes() error {
	if !hasViolation(ir.ValidateEnvelope(runtimeWithoutRuntimeID()), "MISSING_FIELD", "RuntimeID") {
		return errors.New("RuntimeScope missing RuntimeID was not rejected")
	}
	if !hasViolation(ir.ValidateEnvelope(taskWithRuntimeID()), "FORBIDDEN_FIELD", "RuntimeID") {
		return errors.New("TaskScope RuntimeID was not rejected")
	}
	return nil
}
