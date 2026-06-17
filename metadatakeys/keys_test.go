package metadatakeys

import "testing"

func TestKeysAreUnique(t *testing.T) {
	seen := map[Key]bool{}
	for _, key := range All() {
		if key == "" {
			t.Fatal("blank metadata key")
		}
		if seen[key] {
			t.Fatalf("duplicate metadata key %q", key)
		}
		seen[key] = true
	}
}

func TestStoredKeyCompatibility(t *testing.T) {
	tests := map[Key]string{
		WorkspaceID:                  "workspace_id",
		TaskID:                       "task_id",
		RunID:                        "run_id",
		RuntimeLeaseID:               "runtime_lease_id",
		RuntimeCapabilityFingerprint: "runtime_capability_fingerprint",
		ProgressMessageCode:          "riido_progress_message_code",
		ThreadProgressSeq:            "thread_progress_seq",
		AssignmentRecovery:           "recovery",
		HTTPRoute:                    "http.route",
		HTTPResponseStatusCode:       "http.response.status_code",
		HTTPStatusCode:               "http.status_code",
		AWSOperation:                 "aws.operation",
		RiidoTraceSurface:            "riido.trace.surface",
		RiidoStoreOperation:          "riido.store.operation",
		RiidoTaskContextOperation:    "riido.task_context.operation",
		RiidoPollAction:              "riido.poll.action",
		AssignmentResultStatus:       "assignment_result_status",
		AssignmentFailureCategory:    "assignment_failure_category",
		AssignmentEventKey:           "assignment_event_key",
	}
	for key, want := range tests {
		if got := key.String(); got != want {
			t.Fatalf("%s = %q, want %q", key, got, want)
		}
	}
}
