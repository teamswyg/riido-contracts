package ir

import (
	"testing"
	"time"
)

func TestValidateEnvelope_RunScopeValid(t *testing.T) {
	assertNoEnvelopeViolations(t, validRunScopeEvent())
}

func TestValidateEnvelope_RunScopeMissingRuntimeID(t *testing.T) {
	e := validRunScopeEvent()
	e.RuntimeID = ""
	v := ValidateEnvelope(e)
	if !hasViolation(v, "MISSING_FIELD", "RuntimeID") {
		t.Fatalf("expected MISSING_FIELD RuntimeID, got %+v", v)
	}
}

func TestValidateEnvelope_TaskScopeNoRuntimeRequired(t *testing.T) {
	e := validTaskScopeEvent(EventTaskQueued)
	assertNoEnvelopeViolations(t, e)
}

func TestValidateEnvelope_TaskScopeForbidsRuntimeID(t *testing.T) {
	e := validTaskScopeEvent(EventTaskCreated)
	e.RuntimeID = "rt_X"
	v := ValidateEnvelope(e)
	if !hasViolation(v, "FORBIDDEN_FIELD", "RuntimeID") {
		t.Fatalf("expected FORBIDDEN_FIELD RuntimeID for TaskScope, got %+v", v)
	}
}

func TestValidateEnvelope_UnknownScopeRejected(t *testing.T) {
	e := validRunScopeEvent()
	e.Scope = "bogus-scope"
	v := ValidateEnvelope(e)
	if !hasViolation(v, "UNKNOWN_SCOPE", "Scope") {
		t.Fatalf("expected UNKNOWN_SCOPE, got %+v", v)
	}
}

func TestValidateEnvelope_SystemScopeNoTask(t *testing.T) {
	e := CanonicalEvent{
		EventID:             "ev_sys",
		OccurredAt:          time.Date(2026, 5, 19, 12, 0, 0, 0, time.UTC),
		EventSchemaVersion:  1,
		Scope:               EventScopeSystem,
		Type:                EventPolicyBundleLoaded,
		ActorKind:           ActorSystem,
		RiidoDaemonVersion:  "0.1.0",
		PolicyBundleVersion: "pb-2",
	}
	assertNoEnvelopeViolations(t, e)
}
