package ir

import "testing"

func TestValidateEnvelope_FakePlaceholderBanned(t *testing.T) {
	for _, sentinel := range []string{"unknown", "UNKNOWN", "none", "pending", "tbd", "n/a", " - "} {
		t.Run(sentinel, func(t *testing.T) {
			e := validRunScopeEvent()
			e.RuntimeID = sentinel
			v := ValidateEnvelope(e)
			if !hasViolation(v, "FAKE_PLACEHOLDER", "RuntimeID") {
				t.Fatalf("expected FAKE_PLACEHOLDER for RuntimeID=%q, got %+v", sentinel, v)
			}
		})
	}
}
