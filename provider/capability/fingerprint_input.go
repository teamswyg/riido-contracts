package capability

// CapabilityFingerprintInput is the deterministic input to
// ComputeCapabilityFingerprint.
//
// The fingerprint captures a runtime capability snapshot and binds it to the
// PolicyBundleVersion used for eligibility. It intentionally excludes
// NativeConfigVersion because NCV is task execution context, not runtime
// capability.
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
