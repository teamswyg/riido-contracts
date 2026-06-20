package main

import (
	"errors"
	"fmt"
	"strings"
)

func verifyQuestions(questions []question) error {
	seen := map[string]bool{}
	for _, q := range questions {
		if !filled(q.ID, q.Area, q.Status, q.Question, q.CurrentStance) {
			return errors.New("question fields are incomplete")
		}
		if seen[q.ID] {
			return fmt.Errorf("duplicate question id %q", q.ID)
		}
		seen[q.ID] = true
		if q.Status != "open" && q.Status != "resolved" {
			return fmt.Errorf("unsupported question status %q", q.Status)
		}
		if q.Status == "open" && strings.TrimSpace(q.NextArtifact) == "" {
			return fmt.Errorf("open question %q requires next_artifact", q.ID)
		}
	}
	return nil
}

func filled(values ...string) bool {
	for _, value := range values {
		if strings.TrimSpace(value) == "" {
			return false
		}
	}
	return true
}
