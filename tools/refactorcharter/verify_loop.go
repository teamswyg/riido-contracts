package main

import "strings"

func completeLoop(loop evidenceLoop) bool {
	return strings.TrimSpace(loop.Observation) != "" &&
		strings.TrimSpace(loop.Hypothesis) != "" &&
		strings.TrimSpace(loop.Execute) != "" &&
		strings.TrimSpace(loop.Evaluate) != "" &&
		strings.TrimSpace(loop.Retrospective) != ""
}
