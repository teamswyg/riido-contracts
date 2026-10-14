package task

import "testing"

func TestActiveSubset(t *testing.T) {
	if !StateRunning.IsActive() {
		t.Error("Running must be active")
	}
	if !StateValidating.IsActive() {
		t.Error("Validating must be active")
	}
	if StateBlocked.IsActive() {
		t.Error("Blocked must NOT be active (runtime pinning is not enforced while blocked)")
	}
}
