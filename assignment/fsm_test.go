package assignment

import (
	"reflect"
	"strings"
	"testing"
)

func TestGeneratedAssignmentFSM(t *testing.T) {
	fsm := GeneratedAssignmentFSM()
	if fsm.Name() != "assignment" {
		t.Fatalf("Name() = %q", fsm.Name())
	}
	if !fsm.CanTransition(AssignmentStateCodeLeased, AssignmentStateCodeRunning) {
		t.Fatal("leased -> running must be legal")
	}
	if fsm.CanTransition(AssignmentStateCodeCompleted, AssignmentStateCodeRunning) {
		t.Fatal("completed -> running must be illegal")
	}
	got := fsm.NextStates(AssignmentStateCodeRunning)
	want := []AssignmentStateCode{
		AssignmentStateCodeRunning,
		AssignmentStateCodeCancelling,
		AssignmentStateCodeCancelled,
		AssignmentStateCodeCompleted,
		AssignmentStateCodeFailed,
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("NextStates(Running) = %#v, want %#v", got, want)
	}
	if !strings.Contains(fsm.Mermaid(), "Running --> Completed") {
		t.Fatalf("Mermaid() missing Running transition:\n%s", fsm.Mermaid())
	}
}

func TestGeneratedAssignmentFSMServiceProvider(t *testing.T) {
	provider := GeneratedAssignmentFSMServiceProvider()
	if provider.AssignmentFSM().Name() != "assignment" {
		t.Fatalf("provider returned unexpected FSM %q", provider.AssignmentFSM().Name())
	}
}
