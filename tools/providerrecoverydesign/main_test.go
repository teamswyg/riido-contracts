package main

import (
	"encoding/json"
	"os"
	"testing"
)

func loadDesign(t *testing.T) design {
	t.Helper()
	body, err := os.ReadFile("../../docs/20-domain/provider-recovery.riido.json")
	if err != nil {
		t.Fatal(err)
	}
	var d design
	if err := json.Unmarshal(body, &d); err != nil {
		t.Fatal(err)
	}
	return d
}

func TestRecoveryDesignRejectsUnsafeOrUnfinishablePaths(t *testing.T) {
	if err := validate(loadDesign(t)); err != nil {
		t.Fatal(err)
	}
	for _, kind := range []string{"unbounded_retry", "unverified_recovery", "unknown_retry", "dead_end", "conflicting_event"} {
		t.Run(kind, func(t *testing.T) {
			d := loadDesign(t)
			switch kind {
			case "unbounded_retry":
				for i := range d.Transitions {
					if d.Transitions[i].Source == "BACKOFF" && d.Transitions[i].Target == "REPAIRING" {
						d.Transitions[i].BudgetRequired = false
					}
				}
			case "unverified_recovery":
				for i := range d.Transitions {
					if d.Transitions[i].Target == "RECOVERED" {
						d.Transitions[i].VerificationRequired = false
					}
				}
			case "unknown_retry":
				d.UnknownEventTarget = "REPAIRING"
			case "dead_end":
				d.States = append(d.States, "LOST")
			case "conflicting_event":
				d.Transitions = append(d.Transitions, d.Transitions[0])
			}
			if err := validate(d); err == nil {
				t.Fatal("unsafe recovery design accepted")
			}
		})
	}
}

func TestRecoveryReaderDocumentMatchesDesign(t *testing.T) {
	d := loadDesign(t)
	body, err := os.ReadFile("../../" + d.GeneratedDoc)
	if err != nil || string(body) != render(d) {
		t.Fatal("recovery reader document drift")
	}
}
