package assignment

import "testing"

func TestRecoveryCodeStoredCompatibility(t *testing.T) {
	if got := RecoveryFreshStartRefused.String(); got != "fresh_start_refused" {
		t.Fatalf("RecoveryFreshStartRefused = %q", got)
	}
}
