package main

import "errors"

func verifyLoop(loop evidenceLoop) error {
	if loop.Observation == "" || loop.Hypothesis == "" || loop.Execute == "" {
		return errors.New("loop observe/hypothesis/execute fields are required")
	}
	if loop.Evaluate == "" || loop.Retrospective == "" {
		return errors.New("loop evaluate/retrospective fields are required")
	}
	return nil
}
