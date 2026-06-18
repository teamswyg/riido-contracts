package assignment

const (
	EventAssignmentQueued       = "assignment_queued"
	EventAssignmentLeased       = "assignment_leased"
	EventAssignmentReady        = "assignment_ready"
	EventAssignmentRunning      = "assignment_running"
	EventAssignmentCancelling   = "assignment_cancelling"
	EventAssignmentCancelled    = "assignment_cancelled"
	EventAssignmentCompleted    = "assignment_completed"
	EventAssignmentFailed       = "assignment_failed"
	EventAssignmentStateUpdated = "assignment_state_updated"
	EventRiidoLog               = "riido_log"
	EventProviderSessionPinned  = "provider_session_pinned"
	EventProviderLog            = "provider_log"
	EventProviderWarning        = "provider_warning"
	EventProviderError          = "provider_error"
)

func AllTaskEventTypes() []string {
	return []string{
		EventAssignmentQueued,
		EventAssignmentLeased,
		EventAssignmentReady,
		EventAssignmentRunning,
		EventAssignmentCancelling,
		EventAssignmentCancelled,
		EventAssignmentCompleted,
		EventAssignmentFailed,
		EventAssignmentStateUpdated,
		EventRiidoLog,
		EventProviderSessionPinned,
		EventProviderLog,
		EventProviderWarning,
		EventProviderError,
	}
}
