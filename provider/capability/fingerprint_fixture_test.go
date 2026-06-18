package capability

func fingerprintFixture() CapabilityFingerprintInput {
	return CapabilityFingerprintInput{
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
}
