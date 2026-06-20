package main

import "errors"

func verifyLoop(loop evidenceLoop) error {
	if !filled(
		loop.Observation,
		loop.Hypothesis,
		loop.Execute,
		loop.Evaluate,
		loop.Retrospective,
	) {
		return errors.New("complete evidence loop is required")
	}
	return nil
}
