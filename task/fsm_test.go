package task

import (
	"reflect"
	"strings"
	"testing"

	"github.com/teamswyg/riido-contracts/ir"
)

func TestGeneratedTaskFSM(t *testing.T) {
	fsm := GeneratedTaskFSM()
	if fsm.Name() != "task" {
		t.Fatalf("Name() = %q", fsm.Name())
	}
	if fsm.TypeUnion() != TaskFSMTypeUnionTaskLifecycleFSM {
		t.Fatalf("TypeUnion() = %q", fsm.TypeUnion())
	}
	if got, want := fsm.StartStates(), []TaskStateCode{TaskStateCodeCreated}; !reflect.DeepEqual(got, want) {
		t.Fatalf("StartStates() = %#v, want %#v", got, want)
	}
	wantEnd := []TaskStateCode{TaskStateCodeCompleted, TaskStateCodeFailed, TaskStateCodeCancelled, TaskStateCodeTimedOut}
	if got := fsm.EndStates(); !reflect.DeepEqual(got, wantEnd) {
		t.Fatalf("EndStates() = %#v, want %#v", got, wantEnd)
	}
	if fsm.PointKind(TaskStateCodeCreated) != TaskFSMPointStart {
		t.Fatalf("Created point kind = %d", fsm.PointKind(TaskStateCodeCreated))
	}
	if fsm.PointKind(TaskStateCodeCompleted) != TaskFSMPointEnd {
		t.Fatalf("Completed point kind = %d", fsm.PointKind(TaskStateCodeCompleted))
	}
	if fsm.PointKind(TaskStateCodeRunning) != TaskFSMPointIntermediate {
		t.Fatalf("Running point kind = %d", fsm.PointKind(TaskStateCodeRunning))
	}
	if fsm.PointKind(TaskStateCodeUnknown) != TaskFSMPointUnknown {
		t.Fatalf("Unknown point kind = %d", fsm.PointKind(TaskStateCodeUnknown))
	}
	if !fsm.CanTransition(TaskStateCodeRunning, TaskStateCodeValidating, ir.EventTypeCodeRunReportedDone) {
		t.Fatal("Running --RunReportedDone--> Validating must be legal")
	}
	if fsm.CanTransition(TaskStateCodeCreated, TaskStateCodeRunning, ir.EventTypeCodeRunStarted) {
		t.Fatal("Created --RunStarted--> Running must be illegal")
	}
	got := fsm.NextStates(TaskStateCodeHumanReview, ir.EventTypeCodeHumanRejected)
	want := []TaskStateCode{TaskStateCodeReworkQueued, TaskStateCodeCancelled}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("NextStates(HumanReview, HumanRejected) = %#v, want %#v", got, want)
	}
	if next := fsm.NextStates(TaskStateCodeCompleted, ir.EventTypeCodeTaskFailed); next != nil {
		t.Fatalf("terminal state must not produce next states: %#v", next)
	}
	if !strings.Contains(fsm.Mermaid(), "Running --> Validating : RunReportedDone") {
		t.Fatalf("Mermaid() missing Running transition:\n%s", fsm.Mermaid())
	}
}

func TestGeneratedTaskFSMServiceProvider(t *testing.T) {
	provider := GeneratedTaskFSMServiceProvider()
	if provider.TaskFSM().Name() != "task" {
		t.Fatalf("provider returned unexpected FSM %q", provider.TaskFSM().Name())
	}
}
