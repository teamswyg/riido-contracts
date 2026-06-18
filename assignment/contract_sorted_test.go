package assignment

import "slices"

func sortedAssignmentStateSet(values map[AssignmentState]struct{}) []AssignmentState {
	keys := make([]AssignmentState, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	return keys
}

func sortedPollActionSet(values map[PollAction]struct{}) []PollAction {
	keys := make([]PollAction, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	return keys
}

func sortedStringSet(values map[string]struct{}) []string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	return keys
}
