package capability

import "encoding/json"

type fingerprintPayload struct {
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
}

func newFingerprintPayload(in CapabilityFingerprintInput) (fingerprintPayload, error) {
	flagsCanon, err := marshalSortedMap(in.ImportantSurfaceFlags)
	if err != nil {
		return fingerprintPayload{}, err
	}
	return fingerprintPayload{
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
	}, nil
}
