// Package capability owns the C3 Provider Capability domain: the static
// model of "what a provider can do".
//
// What this package does NOT own:
//   - The Provider port and adapter execution lifecycle (process / session /
//     run / ACL) → daemon provider runtime packages (C4).
//   - Scheduling decisions, lease management → daemon scheduling packages (C5).
//   - The flock / DB lease primitives themselves → daemon lock/lease packages
//     (C9).
//
// Dependency direction: capability is a leaf package. Other domain packages
// import it; it imports none of them.
package capability

// RuntimeID is the stable identifier of a registered runtime slot.
// It persists across re-detection. See the Provider Capability generated reader
// and package invariants.
//
// NOTE: a single ambiguous identifier merging runtime slot + vendor family
// + protocol selection is intentionally NOT used. RuntimeID is the runtime
// slot, ProviderKind is the vendor family, ProtocolKind is the
// adapter/protocol selection — three distinct concepts.
// See the Provider Capability invariant that ProtocolKind is the adapter key
// and the doc-consistency-audit-2026-05-19.md §9 banned-name rule.
type RuntimeID string

// CapabilityFingerprint is the SHA-256 hex of the effective capability
// snapshot. "Effective" means the input includes PolicyBundleVersion — same
// binary with a different runtime eligibility policy yields a different
// fingerprint. NativeConfigVersion is intentionally excluded because it is
// task/run execution context, not runtime capability. See
// provider-capability generated reader's fingerprint invariant.
type CapabilityFingerprint string

// DetectedFingerprint is the hash of the provider binary itself (checksum +
// CLI banner). It changes only when the binary changes; a change here
// triggers capability re-detection (which may then yield a new
// CapabilityFingerprint).
type DetectedFingerprint string

// ProviderKind is the vendor family identifier.
// Examples: "claude", "codex", "claude-wrapper".
type ProviderKind string

// ProtocolKind is the primary key for adapter selection.
// See provider-capability invariant anchors: ProviderKind alone does NOT
// determine the adapter — same ProviderKind may host multiple ProtocolKinds
// (e.g., codex-exec-jsonl vs codex-app-server).
type ProtocolKind string

const (
	ProtocolClaudeStreamJSON        ProtocolKind = "claude-stream-json"
	ProtocolCodexExecJSONL          ProtocolKind = "codex-exec-jsonl"
	ProtocolCodexAppServer          ProtocolKind = "codex-app-server"
	ProtocolClaudeCompatibleWrapper ProtocolKind = "claude-compatible-wrapper"
	ProtocolOpenClawAgentJSON       ProtocolKind = "openclaw-agent-json"
	ProtocolCursorAgentStreamJSON   ProtocolKind = "cursor-agent-stream-json"
)
