package capability

// EventStreamFormat is the provider-neutral classification of structured
// event streams. Provider-specific flag names (Claude --output-format
// stream-json, Codex exec --json) all map onto these values via the
// adapter ACL.
type EventStreamFormat string

const (
	EventStreamFormatUnknown              EventStreamFormat = "unknown"
	EventStreamFormatTextOnly             EventStreamFormat = "text-only"
	EventStreamFormatNDJSON               EventStreamFormat = "ndjson"
	EventStreamFormatJSONRPCNotifications EventStreamFormat = "json-rpc-notifications"
)

// ProtocolMaturity is the official maturity classification of a protocol.
// Example: Codex app-server is "experimental" per the official CLI reference.
//
// This value is independent of CompatibilityStatus — even a "stable" protocol
// can be Degraded (probe partial) and an "experimental" protocol can be
// Supported once the adapter has full fixture coverage and opt-in is enabled.
type ProtocolMaturity string

const (
	ProtocolMaturityUnknown      ProtocolMaturity = "unknown"
	ProtocolMaturityStable       ProtocolMaturity = "stable"
	ProtocolMaturityExperimental ProtocolMaturity = "experimental"
	ProtocolMaturityDeprecated   ProtocolMaturity = "deprecated"
)

// CompatibilityStatus is the SUMMARY status of a ProviderCapability.
// Priority order (highest first): Blocked > Experimental > Degraded > Supported.
//
// Auxiliary fields on ProviderCapability (ProtocolMaturity, MissingCapabilities,
// BlockedReasons, DegradedReasons, RequiresExperimentalOptIn) carry the
// reasons. Readers MUST consult them — relying on the summary alone loses
// information about *why* a status was assigned. See provider-capability
// invariant anchors.
type CompatibilityStatus string

const (
	CompatSupported    CompatibilityStatus = "supported"
	CompatDegraded     CompatibilityStatus = "degraded"
	CompatExperimental CompatibilityStatus = "experimental"
	CompatBlocked      CompatibilityStatus = "blocked"
)

// CapabilityName is a stable identifier for a capability surface flag.
// Used in ProviderCapability.MissingCapabilities and probe diagnostics.
type CapabilityName string

// CompatibilityReason records why a (Blocked|Degraded) state assignment
// occurred. Multiple reasons may accumulate on a single ProviderCapability.
type CompatibilityReason struct {
	Code    string // e.g. "POLICY_INCOMPATIBLE", "ADAPTER_REGRESSION_FAILED", "PROBE_PARTIAL", "MIN_VERSION"
	Subject string
	Detail  string
}
