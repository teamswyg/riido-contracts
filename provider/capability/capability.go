package capability

import "time"

// ProviderCapability is the immutable snapshot of one runtime's capabilities.
//
// A new instance is produced on each detection. When CapabilityFingerprint
// changes, leases that pinned to the previous fingerprint are invalidated.
type ProviderCapability struct {
	// identity
	RuntimeID    RuntimeID
	ProviderKind ProviderKind
	ProtocolKind ProtocolKind

	// adapter / protocol
	AdapterID       string
	AdapterVersion  string
	ProtocolVersion string

	// discovery
	ExecutablePath        string
	Argv0                 string
	DetectedVersion       string
	DetectedFingerprint   DetectedFingerprint
	CapabilityFingerprint CapabilityFingerprint
	DiscoveredAt          time.Time

	// provider-neutral surface flags (event stream)
	SupportsStructuredEventStream bool
	EventStreamFormat             EventStreamFormat
	SupportsPartialDeltas         bool

	// provider-neutral surface flags (session / resume)
	SupportsResume       bool
	SupportsSessionID    bool
	SupportsSessionPin   bool
	SupportsSystemPrompt bool
	SupportsMaxTurns     bool

	// provider-neutral surface flags (tool / file events)
	SupportsToolEvents   bool
	SupportsFileEvents   bool
	SupportsUsageMetrics bool

	// provider-neutral surface flags (safety / approval)
	SupportsPermissionControl     bool
	ExposesUnsafePermissionBypass bool // RISK SIGNAL — see invariant 4 below
	SupportsApprovalProtocol      bool
	SupportsSandbox               bool
	SupportsManagedSettings       bool

	// provider-neutral surface flags (extensions)
	SupportsHookEvents      bool
	SupportsMCP             bool
	SupportsWorktree        bool
	SupportsJSONSchemaTools bool

	// safety-surface defaults
	DefaultSandboxMode    string
	DefaultApprovalPolicy string
	HasNetworkOffDefault  bool

	// compatibility envelope (summary + auxiliary signals)
	//
	// Readers must consult BOTH CompatibilityStatus and the auxiliary fields
	// below; the summary alone loses information.
	CompatibilityStatus       CompatibilityStatus
	ProtocolMaturity          ProtocolMaturity
	RequiresExperimentalOptIn bool
	MissingCapabilities       []CapabilityName
	BlockedReasons            []CompatibilityReason
	DegradedReasons           []CompatibilityReason
	MinSupportedVersion       string
	MaxTestedVersion          string

	// raw bag — unknown surfaces preserved (adapter ACL "unknown" jar)
	Unknown map[string]any
}

// INVARIANTS (mirror provider-capability.md §0):
//
//  1. ProtocolKind is the primary key for adapter selection. Same ProviderKind
//     can host multiple ProtocolKinds; matching on ProviderKind alone is
//     forbidden.
//
//  2. DetectedVersion is a raw signal, never a branch condition. Comparing
//     versions like ">= 2.1.42" to gate behavior is forbidden — use surface
//     flags and ProtocolKind / ProtocolVersion instead.
//
//  3. CompatibilityStatus is a SUMMARY of four inputs (detection × maturity ×
//     policy compatibility × adapter test). Readers MUST consult auxiliary
//     fields for the reason.
//
//  4. ExposesUnsafePermissionBypass = true is a RISK SIGNAL, not a license to
//     use the bypass. The final use decision is made by the C7 Security/Policy
//     gate, not by adapter or scheduler code.
//
//  5. Event stream format identifiers are provider-neutral. Claude's
//     "stream-json" and Codex's "exec --json" are both EventStreamFormatNDJSON.
//     Vendor-specific flag names live on a separate (future) surface struct.
