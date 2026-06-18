package task

import (
	"encoding/json"
	"testing"
)

func TestTaskStateGeneratedCodeRoundTripAndStorageJSON(t *testing.T) {
	if got := StateQueued.Code(); got != TaskStateCodeQueued {
		t.Fatalf("StateQueued.Code() = %v, want %v", got, TaskStateCodeQueued)
	}
	if got := TaskStateCodeQueued.StringValue(); got != TaskStateStringQueued {
		t.Fatalf("TaskStateCodeQueued.StringValue() = %q, want %q", got, TaskStateStringQueued)
	}
	if got := TaskStateCodeQueued.TaskState(); got != StateQueued {
		t.Fatalf("TaskStateCodeQueued.TaskState() = %q, want %q", got, StateQueued)
	}
	if ParseTaskStateCode("not-a-state") != TaskStateCodeUnknown {
		t.Fatal("unknown task state must parse to TaskStateCodeUnknown")
	}
	body, err := json.Marshal(struct {
		State TaskState `json:"state"`
	}{State: StateQueued})
	if err != nil {
		t.Fatal(err)
	}
	if got, want := string(body), `{"state":"Queued"}`; got != want {
		t.Fatalf("storage JSON = %s, want %s", got, want)
	}
}
