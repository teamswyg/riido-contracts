package ir

import "testing"

func TestReducerErrorMessage(t *testing.T) {
	var err error = &ReducerError{
		Code:    "IR_REDUCER_INCOMPAT",
		EventID: "ev_1",
		Detail:  "no dispatch for (TaskQueued, 999)",
	}
	if err.Error() == "" {
		t.Fatal("expected non-empty error message")
	}
}
