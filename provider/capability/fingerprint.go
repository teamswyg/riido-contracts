package capability

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
)

// CapabilityFingerprintInput is the deterministic input to
// ComputeCapabilityFingerprint.
//
// The fingerprint captures a RUNTIME CAPABILITY SNAPSHOT — "what can this
// runtime do" — and binds the runtime to a specific PolicyBundleVersion
// (used to evaluate runtime eligibility). It does NOT include the task-level
// NativeConfigVersion, because:
//
//  1. TaskClaimed is RunScope and must carry CapabilityFingerprint, but at
//     claim time the workspace has not been prepared yet — NativeConfigVersion
//     does not exist. If NCV were a fingerprint input, claim itself would be
//     impossible.
//  2. NCV is task-specific execution context, not runtime capability. Runtime
//     pinning binds (RuntimeID, CapabilityFingerprint); execution pinning
//     additionally binds (PolicyBundleVersion, NativeConfigVersion) but those
//     are tracked separately on the run.
//
// See provider-capability.md §2.1 / §2.2 (#27i boundary patch).
//
// ImportantSurfaceFlags is the set of capability flags whose change must
// invalidate leases — see the list in provider-capability.md §2.1.
type CapabilityFingerprintInput struct {
	ProviderKind          ProviderKind        `json:"providerKind"`
	ProtocolKind          ProtocolKind        `json:"protocolKind"`
	ProviderVersion       string              `json:"providerVersion"`
	DetectedFingerprint   DetectedFingerprint `json:"detectedFingerprint"`
	AdapterID             string              `json:"adapterID"`
	AdapterVersion        string              `json:"adapterVersion"`
	ProtocolVersion       string              `json:"protocolVersion"`
	DefaultSandboxMode    string              `json:"defaultSandboxMode"`
	DefaultApprovalPolicy string              `json:"defaultApprovalPolicy"`
	PolicyBundleVersion   string              `json:"policyBundleVersion"`
	ImportantSurfaceFlags map[string]any      `json:"importantSurfaceFlags"`
}

// ComputeCapabilityFingerprint produces the SHA-256 hex of the canonical
// JSON serialization of input. Same input from any daemon must yield the
// same fingerprint (deterministic).
//
// MVP canonicalization: ImportantSurfaceFlags is serialized as a sorted-key
// array of {k,v}; the top-level fields are emitted in struct order via
// json.Marshal. A stricter standard (e.g., RFC 8785 JCS) can replace this
// without changing the public surface.
func ComputeCapabilityFingerprint(in CapabilityFingerprintInput) (CapabilityFingerprint, error) {
	flagsCanon, err := marshalSortedMap(in.ImportantSurfaceFlags)
	if err != nil {
		return "", fmt.Errorf("canonicalize surface flags: %w", err)
	}
	payload := struct {
		ProviderKind          ProviderKind        `json:"providerKind"`
		ProtocolKind          ProtocolKind        `json:"protocolKind"`
		ProviderVersion       string              `json:"providerVersion"`
		DetectedFingerprint   DetectedFingerprint `json:"detectedFingerprint"`
		AdapterID             string              `json:"adapterID"`
		AdapterVersion        string              `json:"adapterVersion"`
		ProtocolVersion       string              `json:"protocolVersion"`
		DefaultSandboxMode    string              `json:"defaultSandboxMode"`
		DefaultApprovalPolicy string              `json:"defaultApprovalPolicy"`
		PolicyBundleVersion   string              `json:"policyBundleVersion"`
		ImportantSurfaceFlags json.RawMessage     `json:"importantSurfaceFlags"`
	}{
		ProviderKind:          in.ProviderKind,
		ProtocolKind:          in.ProtocolKind,
		ProviderVersion:       in.ProviderVersion,
		DetectedFingerprint:   in.DetectedFingerprint,
		AdapterID:             in.AdapterID,
		AdapterVersion:        in.AdapterVersion,
		ProtocolVersion:       in.ProtocolVersion,
		DefaultSandboxMode:    in.DefaultSandboxMode,
		DefaultApprovalPolicy: in.DefaultApprovalPolicy,
		PolicyBundleVersion:   in.PolicyBundleVersion,
		ImportantSurfaceFlags: flagsCanon,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("marshal payload: %w", err)
	}
	sum := sha256.Sum256(b)
	return CapabilityFingerprint(hex.EncodeToString(sum[:])), nil
}

func marshalSortedMap(m map[string]any) ([]byte, error) {
	if m == nil {
		return []byte("[]"), nil
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	type kv struct {
		K string `json:"k"`
		V any    `json:"v"`
	}
	out := make([]kv, 0, len(keys))
	for _, k := range keys {
		out = append(out, kv{K: k, V: m[k]})
	}
	return json.Marshal(out)
}
