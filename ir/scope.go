package ir

// EventScope classifies a CanonicalEvent by which identity fields it must
// carry. See docs/20-domain/ir-schema-versioning.md.
//
// Scope assignment rules are summarized by the generated IR Event Log reader:
//   - SystemScope:  daemon / operator / policy-bundle events. No TaskID/RunID/RuntimeID.
//   - RuntimeScope: runtime slot + capability snapshot events. RuntimeID present
//     (CapabilityFingerprint present once capability is detected).
//     No TaskID/RunID.
//   - TaskScope:    task-level events before claim — TaskCreated/TaskQueued and
//     pre-claim cancellations. TaskID present, no RunID/RuntimeID.
//   - RunScope:     events from a specific (TaskID, RunID). Full runtime fingerprint
//     required (RuntimeID, CapabilityFingerprint, plus 9-axis identifiers).
type EventScope string

const (
	EventScopeSystem  EventScope = "system"
	EventScopeRuntime EventScope = "runtime"
	EventScopeTask    EventScope = "task"
	EventScopeRun     EventScope = "run"
)

// IsValid reports whether s is one of the four declared scopes.
// Unknown values must be rejected by ValidateEnvelope.
func (s EventScope) IsValid() bool {
	switch s {
	case EventScopeSystem, EventScopeRuntime, EventScopeTask, EventScopeRun:
		return true
	default:
		return false
	}
}
