package main

import "errors"

func verifyLoop(loop evidenceLoop) error {
	if loop.Observation == "" || loop.Hypothesis == "" || loop.Execute == "" {
		return errors.New("loop observation, hypothesis, and execute are required")
	}
	if loop.Evaluate == "" || loop.Retrospective == "" {
		return errors.New("loop evaluate and retrospective are required")
	}
	return nil
}
