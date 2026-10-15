package assignment

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestExecutionContextMetadataIsBoundedAndLegacyRemainsAbsent(t *testing.T) {
	if ExecutionContextFromMetadata(map[string]string{"workdir": "/private/customer"}) != nil {
		t.Fatal("legacy path metadata must not invent execution evidence")
	}
	input := ExecutionContext{"repository_clone", "passed", "none", "none"}
	if got := ExecutionContextFromMetadata(input.Metadata()); got == nil || *got != input {
		t.Fatalf("execution evidence lost: %+v", got)
	}
	unsafe := ExecutionContext{"/private/customer", "passed with secret", "token-value", "run arbitrary command"}
	got := ExecutionContextFromMetadata(unsafe.Metadata())
	if got.DiagnosticCode != "unknown" || got.VerificationStatus != "unknown" || got.NextAction != "review_result" {
		t.Fatal("untrusted diagnostics must remain unknown and require review")
	}
	body, _ := json.Marshal(got)
	if strings.Contains(string(body), "private") || strings.Contains(string(body), "secret") || strings.Contains(string(body), "token") {
		t.Fatal("local or free-form data crossed shared diagnostic boundary")
	}
}
