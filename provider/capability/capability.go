package capability

import "time"

// ProviderCapability is one immutable runtime capability snapshot.
type ProviderCapability struct {
	RuntimeID    RuntimeID
	ProviderKind ProviderKind
	ProtocolKind ProtocolKind

	AdapterID       string
	AdapterVersion  string
	ProtocolVersion string

	ExecutablePath        string
	Argv0                 string
	DetectedVersion       string
	DetectedFingerprint   DetectedFingerprint
	CapabilityFingerprint CapabilityFingerprint
	DiscoveredAt          time.Time

	SupportsStructuredEventStream bool
	EventStreamFormat             EventStreamFormat
	SupportsPartialDeltas         bool

	SupportsResume       bool
	SupportsSessionID    bool
	SupportsSessionPin   bool
	SupportsSystemPrompt bool
	SupportsMaxTurns     bool

	SupportsToolEvents   bool
	SupportsFileEvents   bool
	SupportsUsageMetrics bool

	SupportsPermissionControl     bool
	ExposesUnsafePermissionBypass bool
	SupportsApprovalProtocol      bool
	SupportsSandbox               bool
	SupportsManagedSettings       bool

	SupportsHookEvents      bool
	SupportsMCP             bool
	SupportsWorktree        bool
	SupportsJSONSchemaTools bool

	DefaultSandboxMode    string
	DefaultApprovalPolicy string
	HasNetworkOffDefault  bool

	CompatibilityStatus       CompatibilityStatus
	ProtocolMaturity          ProtocolMaturity
	RequiresExperimentalOptIn bool
	MissingCapabilities       []CapabilityName
	BlockedReasons            []CompatibilityReason
	DegradedReasons           []CompatibilityReason
	MinSupportedVersion       string
	MaxTestedVersion          string

	Unknown map[string]any
}
