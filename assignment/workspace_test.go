package assignment

import (
	"encoding/json"
	"testing"
)

func TestAssignmentWorkspaceIsAdditiveAndRoundTrips(t *testing.T) {
	for _, body := range []string{`{"assignment_id":"one"}`, `{"assignment_id":"one","workspace_id":"workspace-a"}`} {
		var value Assignment
		if err := json.Unmarshal([]byte(body), &value); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(value)
		if err != nil {
			t.Fatal(err)
		}
		var wire map[string]any
		if err := json.Unmarshal(encoded, &wire); err != nil {
			t.Fatal(err)
		}
		if value.WorkspaceID == "" {
			if _, ok := wire["workspace_id"]; ok {
				t.Fatal("legacy payload gained workspace")
			}
		} else if wire["workspace_id"] != "workspace-a" {
			t.Fatal("workspace was not preserved")
		}
	}
}
