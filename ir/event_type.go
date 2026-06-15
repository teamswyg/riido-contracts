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
//
// The concrete constants, transition classification, and iota-backed
// EventTypeCode mapping are generated from enumgen/enums.lisp.
type EventType string
