package assignment

import "testing"

func assertTaskEventsMatchContract(t *testing.T, contract executableContract) {
	t.Helper()
	remaining := map[string]struct{}{}
	for _, event := range AllTaskEventTypes() {
		remaining[event] = struct{}{}
	}
	for _, event := range contract.TaskEvents {
		if _, ok := remaining[event.Value]; !ok {
			t.Fatalf("contract task event %q is missing from package constants", event.Value)
		}
		delete(remaining, event.Value)
	}
	if len(remaining) != 0 {
		t.Fatalf("package task events missing from contract: %v", sortedStringSet(remaining))
	}
}
