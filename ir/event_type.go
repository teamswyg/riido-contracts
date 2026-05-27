// Package ir owns the C2 IR Event Log domain types: CanonicalEvent, EventType,
// ActorKind, and the pure Reducer contract.
//
// The catalog and dispatch rules implemented here are the code-level reflection
// of docs/20-domain/ir-event-log.md. The CanonicalEvent field requirements
// follow docs/20-domain/ir-schema-versioning.md §2.
//
// PURITY INVARIANT (ir-event-log.md §5.0): Reducer MUST NOT append events.
// This package contains NO writer / appender / I/O dependency.
package ir

// EventType is one of the catalog entries from ir-event-log.md §3.
//
// Categories:
//
//	A — Task lifecycle transitions (all transition events)
//	B — Runtime registry / capability lifecycle
//	C — Provider raw → canonical (adapter ACL output, non-transition)
//	D — Validation (ValidationPassed/Failed counted under A)
//	E — Workspace / config injection (non-transition)
//	F — Security / policy (non-transition)
//	G — Upgrade / runtime change
//	H — Administrative / audit (non-transition)
type EventType string

// Category A — Task lifecycle transitions.
const (
	EventTaskCreated            EventType = "TaskCreated"
	EventTaskQueued             EventType = "TaskQueued"
	EventTaskClaimed            EventType = "TaskClaimed"
	EventWorkdirPreparing       EventType = "WorkdirPreparing"
	EventRuntimePinned          EventType = "RuntimePinned"
	EventRunStarted             EventType = "RunStarted"
	EventInputRequested         EventType = "InputRequested"
	EventInputProvided          EventType = "InputProvided"
	EventBlockerRaised          EventType = "BlockerRaised"
	EventBlockerResolved        EventType = "BlockerResolved"
	EventBlockerResolvedRequeue EventType = "BlockerResolvedRequeue"
	EventRunReportedDone        EventType = "RunReportedDone"
	EventValidationPassed       EventType = "ValidationPassed"
	EventValidationFailed       EventType = "ValidationFailed"
	EventReviewRequested        EventType = "ReviewRequested"
	EventAutoApproved           EventType = "AutoApproved"
	EventHumanApproved          EventType = "HumanApproved"
	EventHumanRejected          EventType = "HumanRejected"
	EventReworkAccepted         EventType = "ReworkAccepted"
	EventTaskCancelled          EventType = "TaskCancelled"
	EventTaskTimedOut           EventType = "TaskTimedOut"
	EventRuntimePinViolated     EventType = "RuntimePinViolated"
	EventTaskFailed             EventType = "TaskFailed"
)

// Category B — Runtime registry / capability lifecycle.
const (
	EventRuntimeRegistered         EventType = "RuntimeRegistered"
	EventRuntimeRejected           EventType = "RuntimeRejected"
	EventRuntimeFingerprintChanged EventType = "RuntimeFingerprintChanged"
	EventCapabilityReevaluated     EventType = "CapabilityReevaluated"
	EventLeaseInvalidated          EventType = "LeaseInvalidated"
	EventRuntimeHandshakeOK        EventType = "RuntimeHandshakeOK"
)

// Category C — Provider raw → canonical (adapter ACL output).
const (
	EventTextDelta            EventType = "TextDelta"
	EventReasoningDelta       EventType = "ReasoningDelta"
	EventToolCallStarted      EventType = "ToolCallStarted"
	EventToolCallFinished     EventType = "ToolCallFinished"
	EventFileChanged          EventType = "FileChanged"
	EventCommandStarted       EventType = "CommandStarted"
	EventCommandFinished      EventType = "CommandFinished"
	EventSessionPinned        EventType = "SessionPinned"
	EventApprovalRequested    EventType = "ApprovalRequested"
	EventApprovalResolved     EventType = "ApprovalResolved"
	EventStatusUpdate         EventType = "StatusUpdate"
	EventUsageDelta           EventType = "UsageDelta"
	EventLogLine              EventType = "LogLine"
	EventProviderUnknownEvent EventType = "ProviderUnknownEvent"
)

// Category D — Validation (non-transition members).
const (
	EventValidationStarted      EventType = "ValidationStarted"
	EventValidationRuleExecuted EventType = "ValidationRuleExecuted"
)

// Category E — Workspace / config injection.
const (
	EventWorkdirCreated           EventType = "WorkdirCreated"
	EventNativeConfigInjected     EventType = "NativeConfigInjected"
	EventWorkdirArchived          EventType = "WorkdirArchived"
	EventConfigTemplateReinjected EventType = "ConfigTemplateReinjected"
)

// Category F — Security / policy.
const (
	EventPolicyBundleLoaded      EventType = "PolicyBundleLoaded"
	EventPolicyBundleSwitched    EventType = "PolicyBundleSwitched"
	EventPolicyViolationDetected EventType = "PolicyViolationDetected"
	EventSecretsScopeIssued      EventType = "SecretsScopeIssued"
	EventSecretsScopeRevoked     EventType = "SecretsScopeRevoked"
)

// Category G — Upgrade / runtime change.
const (
	EventUpgradeDetected          EventType = "UpgradeDetected"
	EventUpgradePolicyReevaluated EventType = "UpgradePolicyReevaluated"
	EventDrainStarted             EventType = "DrainStarted"
	EventDrainTimedOut            EventType = "DrainTimedOut"
	EventTaskHandedOff            EventType = "TaskHandedOff"
)

// Category H — Administrative / audit.
const (
	EventCorrection   EventType = "Correction"
	EventOperatorNote EventType = "OperatorNote"
	EventSnapshot     EventType = "Snapshot"
)

// transitionSet is the SSOT for "which EventType is a FSM transition event".
// Mirrors ir-event-log.md §3 "Transition: ✓" column.
//
// Invariant: every transition event MUST carry a non-zero FSMVersion when
// appended (ir-schema-versioning.md §0 invariant 4).
var transitionSet = map[EventType]struct{}{
	EventTaskCreated:            {},
	EventTaskQueued:             {},
	EventTaskClaimed:            {},
	EventWorkdirPreparing:       {},
	EventRuntimePinned:          {},
	EventRunStarted:             {},
	EventInputRequested:         {},
	EventInputProvided:          {},
	EventBlockerRaised:          {},
	EventBlockerResolved:        {},
	EventBlockerResolvedRequeue: {},
	EventRunReportedDone:        {},
	EventValidationPassed:       {},
	EventValidationFailed:       {},
	EventReviewRequested:        {},
	EventAutoApproved:           {},
	EventHumanApproved:          {},
	EventHumanRejected:          {},
	EventReworkAccepted:         {},
	EventTaskCancelled:          {},
	EventTaskTimedOut:           {},
	EventRuntimePinViolated:     {},
	EventTaskFailed:             {},
}

// IsTransition reports whether t is classified as a FSM transition event.
func (t EventType) IsTransition() bool {
	_, ok := transitionSet[t]
	return ok
}
