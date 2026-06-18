package ir

import (
	"testing"
	"time"
)

func validRunScopeEvent() CanonicalEvent {
	return CanonicalEvent{
		EventID:               "ev_1",
		OccurredAt:            time.Date(2026, 5, 19, 12, 0, 0, 0, time.UTC),
		EventSchemaVersion:    1,
		Scope:                 EventScopeRun,
		Type:                  EventTextDelta,
		ActorKind:             ActorAgent,
		ActorID:               "run_1",
		RiidoDaemonVersion:    "0.1.0",
		PolicyBundleVersion:   "pb-1",
		TaskID:                "task_1",
		RunID:                 "run_1",
		RuntimeID:             "rt_1",
		CapabilityFingerprint: "fp_abc",
		ProviderKind:          "claude",
		ProtocolKind:          "claude-stream-json",
		ProviderVersion:       "2.1.128",
		AdapterID:             "claude-stream-json",
		AdapterVersion:        "0.1.0",
		ProtocolVersion:       "stream-json-v1",
		NativeConfigVersion:   "nc_xyz",
	}
}

func hasViolation(vs []EnvelopeViolation, code, field string) bool {
	for _, v := range vs {
		if v.Code == code && v.Field == field {
			return true
		}
	}
	return false
}

func transitionEvent(e CanonicalEvent, eventType EventType) CanonicalEvent {
	e.Type = eventType
	if eventType.IsTransition() {
		e.FSMVersion = 1
	}
	return e
}

func assertNoEnvelopeViolations(t *testing.T, e CanonicalEvent) {
	t.Helper()
	if v := ValidateEnvelope(e); len(v) != 0 {
		t.Fatalf("event should be envelope-valid, got %+v", v)
	}
}
