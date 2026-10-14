package assignment

import (
	"encoding/json"
	"testing"
)

func assertApprovalJSON(t *testing.T, value any, want string) {
	t.Helper()
	data, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if got := string(data); got != want {
		t.Fatalf("json = %s, want %s", got, want)
	}
}
