package assignment

const (
	// ContractSchemaVersion identifies the executable assignment contract
	// fixture in assignment_contract.riido.json.
	ContractSchemaVersion = "riido-ai-server-contract.v1"

	// SchemaVersion is the C10 SaaS assignment API schema version shared by
	// daemon poll/event clients and the control-plane API surface.
	SchemaVersion = "riido-ai-server.v1"

	// RecommendedHeartbeatIntervalSeconds is the daemon-side cadence for
	// active assignment heartbeats.
	RecommendedHeartbeatIntervalSeconds = 5

	// ActiveAssignmentStaleAfterSeconds is the control-plane lease timeout for
	// assignments that are leased, ready, or running.
	ActiveAssignmentStaleAfterSeconds = 20
)
