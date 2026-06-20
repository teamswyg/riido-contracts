// Package task owns the generated C1 Task Lifecycle domain: TaskState enum,
// legal transition matrix, terminal definition, and the FSM SPI generated from
// enumgen/enums.lisp.
//
// What this package does NOT own:
//   - EventType catalog and payload schema → ir (C2).
//   - "Which event is a transition event" classification → ir.
//   - Reducer behavior / dispatch / unknown-pair handling → ir.
//
// Dependency direction: task → ir (read-only). ir does NOT depend on task.
package task

// FSMSchemaVersion is the current version of the generated TaskState
// transition matrix.
const FSMSchemaVersion = 1

// TaskState is one of the 15 lifecycle states generated from enumgen/enums.lisp.
//
// The concrete constants and the iota-backed TaskStateCode mapping are
// generated from enumgen/enums.lisp.
type TaskState string
