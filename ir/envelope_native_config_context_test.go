package ir

import "testing"

func TestValidateEnvelopeWithRunContext_PhaseDependent(t *testing.T) {
	for _, c := range phaseDependentRunContextCases() {
		t.Run(c.name, func(t *testing.T) {
			ev := transitionEvent(validRunScopeEvent(), c.evt)
			if !c.ncvSet {
				ev.NativeConfigVersion = ""
			}
			v := ValidateEnvelopeWithRunContext(ev, RunContext{NativeConfigEstablished: c.established})
			got := hasViolation(v, "MISSING_FIELD", "NativeConfigVersion")
			if got != c.wantNCVMissing {
				t.Fatalf("%s: want NCV missing=%v, got %v. Violations: %+v", c.name, c.wantNCVMissing, got, v)
			}
		})
	}
}

type phaseDependentRunContextCase struct {
	name           string
	evt            EventType
	established    bool
	ncvSet         bool
	wantNCVMissing bool
}

func phaseDependentRunContextCases() []phaseDependentRunContextCase {
	return []phaseDependentRunContextCase{
		{"TaskFailed before NCV established", EventTaskFailed, false, false, false},
		{"TaskFailed after NCV established without NCV", EventTaskFailed, true, false, true},
		{"TaskFailed after NCV established with NCV", EventTaskFailed, true, true, false},
		{"RuntimePinViolated before NCV established", EventRuntimePinViolated, false, false, false},
		{"RuntimePinViolated after NCV established without NCV", EventRuntimePinViolated, true, false, true},
		{"BlockerRaised before NCV established", EventBlockerRaised, false, false, false},
		{"BlockerRaised after NCV established without NCV", EventBlockerRaised, true, false, true},
	}
}
