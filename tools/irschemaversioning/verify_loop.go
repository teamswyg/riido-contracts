package main

import "errors"

func verifyLoop(loop evidenceLoop) error {
	if loop.Observation == "" || loop.Hypothesis == "" || loop.Execute == "" {
		return errors.New("evidence loop is incomplete")
	}
	if loop.Evaluate == "" || loop.Retrospective == "" {
		return errors.New("evidence loop is incomplete")
	}
	return nil
}
