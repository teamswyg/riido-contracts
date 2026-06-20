package main

import "fmt"

func verifyLoop(loop evidenceLoop) error {
	if loop.Observation == "" || loop.Hypothesis == "" || loop.Execute == "" {
		return fmt.Errorf("evidence loop is missing observe/hypothesis/execute")
	}
	if loop.Evaluate == "" || loop.Retrospective == "" {
		return fmt.Errorf("evidence loop is missing evaluate/retrospective")
	}
	return nil
}
