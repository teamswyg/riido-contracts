package assignment

import "testing"

func assertPollActionsMatchContract(t *testing.T, contract executableContract) {
	t.Helper()
	remaining := map[PollAction]struct{}{}
	for _, action := range AllPollActions() {
		remaining[action] = struct{}{}
	}
	for _, action := range contract.PollActions {
		value := PollAction(action.Value)
		if !value.Valid() {
			t.Fatalf("contract poll action %q is missing from package constants", action.Value)
		}
		delete(remaining, value)
	}
	if len(remaining) != 0 {
		t.Fatalf("package poll actions missing from contract: %v", sortedPollActionSet(remaining))
	}
}
