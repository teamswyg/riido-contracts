package main

import (
	"fmt"
	"strings"
)

func verifyRendered(rendered string, m manifest) error {
	for _, want := range m.RequiredMarkers {
		if !strings.Contains(rendered, want) {
			return fmt.Errorf("generated README missing required marker %q", want)
		}
	}
	for _, forbidden := range m.ForbiddenLiterals {
		if strings.Contains(rendered, forbidden) {
			return fmt.Errorf("generated README contains forbidden literal %q", forbidden)
		}
	}
	return nil
}

func completeLoop(loop evidenceLoop) bool {
	return filled(
		loop.Observation,
		loop.Hypothesis,
		loop.Execute,
		loop.Evaluate,
		loop.Retrospective,
	)
}
