// Package task owns the C1 Task Lifecycle domain: TaskState enum, legal
// transition matrix, terminal definition, and the FSM invariants enumerated
// in docs/20-domain/task-lifecycle.md.
//
// What this package does NOT own:
//   - EventType catalog and payload schema → ir (C2).
//   - "Which event is a transition event" classification → ir.
//   - Reducer behavior / dispatch / unknown-pair handling → ir.
//
// Dependency direction: task → ir (read-only). ir does NOT depend on task.
package task

// FSMSchemaVersion is the current version of the TaskState transition
// matrix owned by docs/20-domain/task-lifecycle.md.
const FSMSchemaVersion = 1

// TaskState is one of the 15 lifecycle states defined in
// docs/20-domain/task-lifecycle.md §2.
//
// The concrete constants and the iota-backed TaskStateCode mapping are
// generated from enumgen/enums.lisp.
type TaskState string
