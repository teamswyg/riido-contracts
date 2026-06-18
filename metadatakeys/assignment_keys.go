package metadatakeys

const (
	// AssignmentRecovery keeps the legacy "recovery" storage key for
	// DynamoDB/event-history compatibility.
	AssignmentRecovery        Key = "recovery"
	AssignmentResultStatus    Key = "assignment_result_status"
	AssignmentFailureCategory Key = "assignment_failure_category"
	AssignmentEventKey        Key = "assignment_event_key"
)
