package assignment

import (
	"encoding/json"
	"testing"
)

func TestRequestIdentityRoundTripsAndRemainsOptional(t *testing.T) {
	original := Assignment{ID: "assignment", RequestID: "request", WorkspaceID: "workspace", CreatedBy: "author"}
	raw, err := json.Marshal(original)
	if err != nil {
		t.Fatal(err)
	}
	var decoded Assignment
	if err = json.Unmarshal(raw, &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded.RequestID != original.RequestID || decoded.WorkspaceID != original.WorkspaceID || decoded.CreatedBy != original.CreatedBy {
		t.Fatalf("lost scope: %+v", decoded)
	}
	if err = json.Unmarshal([]byte(`{"assignment_id":"legacy"}`), &decoded); err != nil {
		t.Fatal(err)
	}
	var legacy AssignRequest
	if err = json.Unmarshal([]byte(`{"agent_id":"agent","prompt":"work"}`), &legacy); err != nil || legacy.RequestID != "" {
		t.Fatalf("legacy compatibility: %+v %v", legacy, err)
	}
}
