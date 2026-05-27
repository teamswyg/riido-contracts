package capability

import "testing"

func TestFingerprintDeterministic(t *testing.T) {
	in := CapabilityFingerprintInput{
		ProviderKind:          "claude",
		ProtocolKind:          ProtocolClaudeStreamJSON,
		ProviderVersion:       "2.1.128",
		DetectedFingerprint:   "abc123",
		AdapterID:             "claude-stream-json",
		AdapterVersion:        "0.1.0",
		ProtocolVersion:       "stream-json-v1",
		DefaultSandboxMode:    "workspace-write",
		DefaultApprovalPolicy: "on-request",
		PolicyBundleVersion:   "policies-2026-05-19",
		ImportantSurfaceFlags: map[string]any{
			"SupportsStructuredEventStream": true,
			"EventStreamFormat":             "ndjson",
			"ExposesUnsafePermissionBypass": false,
		},
	}
	fp1, err := ComputeCapabilityFingerprint(in)
	if err != nil {
		t.Fatalf("compute 1: %v", err)
	}
	fp2, err := ComputeCapabilityFingerprint(in)
	if err != nil {
		t.Fatalf("compute 2: %v", err)
	}
	if fp1 != fp2 {
		t.Fatalf("fingerprint not deterministic: %s vs %s", fp1, fp2)
	}
	if len(fp1) != 64 {
		t.Fatalf("expected 64-char SHA-256 hex, got %d", len(fp1))
	}
}

// TestFingerprintMapOrderIndependent verifies that Go's randomized map
// iteration cannot influence the fingerprint — otherwise lease handoff
// across daemons breaks.
func TestFingerprintMapOrderIndependent(t *testing.T) {
	a := CapabilityFingerprintInput{
		ImportantSurfaceFlags: map[string]any{
			"zzz": true,
			"aaa": false,
			"mmm": 42,
		},
	}
	b := CapabilityFingerprintInput{
		ImportantSurfaceFlags: map[string]any{
			"aaa": false,
			"mmm": 42,
			"zzz": true,
		},
	}
	fa, err := ComputeCapabilityFingerprint(a)
	if err != nil {
		t.Fatal(err)
	}
	fb, err := ComputeCapabilityFingerprint(b)
	if err != nil {
		t.Fatal(err)
	}
	if fa != fb {
		t.Fatalf("fingerprint depends on map iteration order — canonicalization broken: %s vs %s", fa, fb)
	}
}

// TestFingerprintChangesOnPolicyBundle verifies the EFFECTIVE-snapshot
// invariant: PolicyBundleVersion is part of the hash input.
func TestFingerprintChangesOnPolicyBundle(t *testing.T) {
	base := CapabilityFingerprintInput{
		ProviderKind:        "claude",
		PolicyBundleVersion: "v1",
	}
	fp1, _ := ComputeCapabilityFingerprint(base)
	base.PolicyBundleVersion = "v2"
	fp2, _ := ComputeCapabilityFingerprint(base)
	if fp1 == fp2 {
		t.Fatal("fingerprint must change when PolicyBundleVersion changes (effective-snapshot invariant)")
	}
}

// TestFingerprintChangesOnSurfaceFlag verifies that an important-surface-flag
// change is enough to invalidate the snapshot. NativeConfigVersion is
// intentionally NOT a fingerprint input (#27i boundary patch) — it's tracked
// separately as execution context.
func TestFingerprintChangesOnSurfaceFlag(t *testing.T) {
	base := CapabilityFingerprintInput{
		ProviderKind: "claude",
		ImportantSurfaceFlags: map[string]any{
			"ExposesUnsafePermissionBypass": false,
		},
	}
	fp1, _ := ComputeCapabilityFingerprint(base)
	base.ImportantSurfaceFlags = map[string]any{
		"ExposesUnsafePermissionBypass": true,
	}
	fp2, _ := ComputeCapabilityFingerprint(base)
	if fp1 == fp2 {
		t.Fatal("fingerprint must change when an important surface flag changes")
	}
}
