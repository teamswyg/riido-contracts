package assignment

import (
	"encoding/json"
	"testing"
)

func TestAssignmentGeneratedCodeRoundTripAndStorageJSON(t *testing.T) {
	if got := AssignmentRunning.Code(); got != AssignmentStateCodeRunning {
		t.Fatalf("AssignmentRunning.Code() = %v, want %v", got, AssignmentStateCodeRunning)
	}
	if got := AssignmentStateCodeRunning.StringValue(); got != AssignmentStateStringRunning {
		t.Fatalf("AssignmentStateCodeRunning.StringValue() = %q, want %q", got, AssignmentStateStringRunning)
	}
	if got := AssignmentStateCodeRunning.AssignmentState(); got != AssignmentRunning {
		t.Fatalf("AssignmentStateCodeRunning.AssignmentState() = %q, want %q", got, AssignmentRunning)
	}
	if got := PollStart.Code(); got != PollActionCodeStart {
		t.Fatalf("PollStart.Code() = %v, want %v", got, PollActionCodeStart)
	}
	if ParseAssignmentStateCode("not-a-state") != AssignmentStateCodeUnknown {
		t.Fatal("unknown assignment state must parse to AssignmentStateCodeUnknown")
	}
	body, err := json.Marshal(struct {
		State  AssignmentState `json:"state"`
		Action PollAction      `json:"action"`
	}{State: AssignmentRunning, Action: PollStart})
	if err != nil {
		t.Fatal(err)
	}
	if got, want := string(body), `{"state":"running","action":"start"}`; got != want {
		t.Fatalf("storage JSON = %s, want %s", got, want)
	}
}
