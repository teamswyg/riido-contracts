package assignment

import (
	"slices"
	"testing"
)

func assertExecutionIdentityContract(t *testing.T, contract executableContract) {
	t.Helper()
	identity := contract.ExecutionIdentity
	keyOrder := []string{"assignment_id", "run_id", "task_id"}
	if got := identity.ExecutionKeyOrder; !slices.Equal(got, keyOrder) {
		t.Fatalf("execution key order drifted: got %v want %v", got, keyOrder)
	}
	resumeOrder := []string{"provider_session_id", "resume_session_id"}
	if got := identity.ResumeSessionKeyOrder; !slices.Equal(got, resumeOrder) {
		t.Fatalf("resume session key order drifted: got %v want %v", got, resumeOrder)
	}
	if got := identity.RunIDDefaultSource; got != "assignment_id" {
		t.Fatalf("run id default source drifted: %q", got)
	}
}
