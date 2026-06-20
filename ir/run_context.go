package ir

// RunContext carries the dynamic facts needed to resolve a PhaseDependent
// EventType's NCV requirement. Owned by the orchestrator / EventIngestor /
// reducer; see ir-schema-versioning and the generated IR Event Log reader.
type RunContext struct {
	// NativeConfigEstablished is true once a NativeConfigInjected event
	// or any later ExecutionBoundOnly event has been appended for the same RunID.
	NativeConfigEstablished bool
}
