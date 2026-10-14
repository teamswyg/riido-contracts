package ir

// ActorKind is the server-decided attribution of an event's origin.
//
// INVARIANT from the generated IR Event Log reader: ActorKind and ActorID are
// determined by the server transition layer, NOT by client/CLI/agent input.
// Provider raw stdout never sets these directly.
type ActorKind string

const (
	ActorAgent  ActorKind = "agent"
	ActorDaemon ActorKind = "daemon"
	ActorHuman  ActorKind = "human"
	ActorSystem ActorKind = "system"
)
