package task

import "testing"

func TestTriggerNamesMatchIR(t *testing.T) {
	for _, tr := range LegalTransitions() {
		if tr.Trigger == "" {
			t.Errorf("empty trigger for transition %s -> %s", tr.From, tr.To)
		}
	}
}

func TestEveryTriggerIsClassifiedAsTransition(t *testing.T) {
	for _, tr := range LegalTransitions() {
		if !tr.Trigger.IsTransition() {
			t.Errorf(
				"drift: %s -> %s uses trigger %q which ir.IsTransition() says is NOT a transition event",
				tr.From, tr.To, tr.Trigger,
			)
		}
	}
}
