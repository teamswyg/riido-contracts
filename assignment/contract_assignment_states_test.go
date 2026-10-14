package assignment

import "testing"

func assertAssignmentStatesMatchContract(t *testing.T, contract executableContract) {
	t.Helper()
	remaining := map[AssignmentState]struct{}{}
	for _, state := range AllAssignmentStates() {
		remaining[state] = struct{}{}
	}
	contractStates := map[AssignmentState]contractState{}
	for _, state := range contract.AssignmentStates {
		value := AssignmentState(state.Value)
		assertContractState(t, state, value, contractStates)
		contractStates[value] = state
		delete(remaining, value)
	}
	if len(remaining) != 0 {
		t.Fatalf("package states missing from contract: %v", sortedAssignmentStateSet(remaining))
	}
	assertTransitionsMatchContract(t, contractStates)
}

func assertContractState(
	t *testing.T,
	state contractState,
	value AssignmentState,
	contractStates map[AssignmentState]contractState,
) {
	t.Helper()
	if !value.Valid() {
		t.Fatalf("contract state %q is missing from package constants", state.Value)
	}
	if _, exists := contractStates[value]; exists {
		t.Fatalf("duplicate contract state %q", state.Value)
	}
	if got := IsTerminal(value); got != state.Terminal {
		t.Fatalf("IsTerminal(%q) = %v, want %v", value, got, state.Terminal)
	}
	if got := IsAgentActive(value); got != state.AgentActive {
		t.Fatalf("IsAgentActive(%q) = %v, want %v", value, got, state.AgentActive)
	}
}

func assertTransitionsMatchContract(t *testing.T, contractStates map[AssignmentState]contractState) {
	t.Helper()
	for _, from := range AllAssignmentStates() {
		allowed := allowedContractTransitions(contractStates[from])
		for _, to := range AllAssignmentStates() {
			_, inContract := allowed[to]
			want := from == to || inContract
			if got := CanTransition(from, to); got != want {
				t.Fatalf("CanTransition(%q,%q) = %v, want %v", from, to, got, want)
			}
		}
	}
}

func allowedContractTransitions(state contractState) map[AssignmentState]struct{} {
	allowed := map[AssignmentState]struct{}{}
	for _, transition := range state.Transitions {
		allowed[AssignmentState(transition)] = struct{}{}
	}
	return allowed
}
