package assignment

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestRecoveryPolicyDefaultsAndExplicitSwitch(t *testing.T) {
	p, err := (*ProviderRecoveryPolicy)(nil).Normalized()
	if err != nil || p.MaxAttempts != 3 || p.MaxElapsedSeconds != 30 || p.MaxSwitches != 0 || len(p.Fallbacks) != 0 {
		t.Fatalf("unsafe legacy defaults: %+v, %v", p, err)
	}
	want := &ProviderRecoveryPolicy{MaxAttempts: 2, MaxElapsedSeconds: 10, MaxSwitches: 1, Fallbacks: []ProviderRecoveryCandidate{{Provider: "codex", ModelID: "gpt-5.4"}}}
	got, err := want.Normalized()
	if err != nil {
		t.Fatal(err)
	}
	got.Fallbacks[0].ModelID = "changed"
	if want.Fallbacks[0].ModelID != "gpt-5.4" {
		t.Fatal("normalization shared the mutable authorization snapshot")
	}
}

func TestRecoveryPolicyRejectsUnboundedOrAmbiguousAuthority(t *testing.T) {
	valid := ProviderRecoveryCandidate{Provider: "claude", ModelID: "sonnet"}
	for _, p := range []ProviderRecoveryPolicy{
		{MaxAttempts: -1},
		{MaxAttempts: 4},
		{MaxElapsedSeconds: -1},
		{MaxElapsedSeconds: 61},
		{MaxSwitches: -1},
		{MaxSwitches: 2},
		{MaxSwitches: 1},
		{Fallbacks: []ProviderRecoveryCandidate{valid}},
		{MaxAttempts: 1, MaxSwitches: 1, Fallbacks: []ProviderRecoveryCandidate{valid}},
		{MaxSwitches: 1, Fallbacks: []ProviderRecoveryCandidate{valid, valid}},
		{MaxSwitches: 1, Fallbacks: []ProviderRecoveryCandidate{{Provider: "unknown", ModelID: "sonnet"}}},
		{MaxSwitches: 1, Fallbacks: []ProviderRecoveryCandidate{{Provider: "claude"}}},
		{MaxSwitches: 1, Fallbacks: []ProviderRecoveryCandidate{{Provider: "claude", ModelID: "model\ncredential"}}},
		{MaxSwitches: 1, Fallbacks: []ProviderRecoveryCandidate{{Provider: "claude", ModelID: strings.Repeat("m", 129)}}},
	} {
		if _, err := p.Normalized(); err == nil {
			t.Errorf("accepted invalid policy: %+v", p)
		}
	}
}

func TestRecoveryPolicyAssignmentJSONAndSnapshot(t *testing.T) {
	p := &ProviderRecoveryPolicy{MaxSwitches: 1, Fallbacks: []ProviderRecoveryCandidate{{Provider: "cursor", ModelID: "auto"}}}
	a := Assignment{RecoveryPolicy: CloneProviderRecoveryPolicy(p)}
	p.Fallbacks[0].Provider = "changed"
	body, err := json.Marshal(a)
	if err != nil {
		t.Fatal(err)
	}
	var got Assignment
	if err := json.Unmarshal(body, &got); err != nil || got.RecoveryPolicy.Fallbacks[0].Provider != "cursor" {
		t.Fatalf("policy snapshot round trip failed: %v", err)
	}
	if CloneProviderRecoveryPolicy(nil) != nil {
		t.Fatal("legacy absent policy must remain absent")
	}
}

func TestRecoveryMetadataIsBoundedAndDoesNotMutateInput(t *testing.T) {
	input := map[string]string{RecoveryStateKey: "recovered", RecoveryReasonKey: "/secret/error", RecoveryProviderKey: "claude", RecoveryCircuitKey: "closed", RecoveryAttemptKey: "9999999", RecoverySwitchCountKey: "-1", RecoveryOutcomeCertaintyKey: "confirmed", "recovery_password": "secret", "run_id": "run-1"}
	got := SanitizeRecoveryMetadata(input)
	if got[RecoveryReasonKey] != "unknown" || got[RecoveryAttemptKey] != "0" || got[RecoverySwitchCountKey] != "0" || got["recovery_password"] != "" || got["run_id"] != "run-1" {
		t.Fatalf("unbounded recovery observation: %+v", got)
	}
	if input[RecoveryReasonKey] != "/secret/error" || input["recovery_password"] != "secret" {
		t.Fatal("input map was mutated")
	}
	if RecoveryContextFromMetadata(nil) != nil || SanitizeRecoveryMetadata(nil) != nil {
		t.Fatal("legacy absence must not manufacture recovery")
	}
	if c := RecoveryContextFromMetadata(RecoveryContext{State: "dispatched", Reason: "none", Provider: "codex", Circuit: "half_open", OutcomeCertainty: "started", Attempt: 2, SwitchCount: 1}.Metadata()); c == nil || c.State != "dispatched" || c.Attempt != 2 {
		t.Fatalf("recovery context round trip: %+v", c)
	}
}
