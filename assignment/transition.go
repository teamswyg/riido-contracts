package assignment

func IsAgentActive(state AssignmentState) bool {
	switch state {
	case AssignmentLeased, AssignmentReady, AssignmentRunning:
		return true
	default:
		return false
	}
}

func IsTerminal(state AssignmentState) bool {
	switch state {
	case AssignmentCancelled, AssignmentCompleted, AssignmentFailed:
		return true
	default:
		return false
	}
}

func CanTransition(from, to AssignmentState) bool {
	if from == to {
		return true
	}
	if IsTerminal(from) {
		return false
	}
	switch from {
	case AssignmentQueued:
		switch to {
		case AssignmentLeased, AssignmentCancelling, AssignmentCancelled, AssignmentFailed:
			return true
		default:
			return false
		}
	case AssignmentLeased:
		switch to {
		case AssignmentReady, AssignmentRunning, AssignmentCancelling, AssignmentCancelled, AssignmentFailed:
			return true
		default:
			return false
		}
	case AssignmentReady:
		switch to {
		case AssignmentRunning, AssignmentCancelling, AssignmentCancelled, AssignmentFailed:
			return true
		default:
			return false
		}
	case AssignmentRunning:
		switch to {
		case AssignmentCompleted, AssignmentFailed, AssignmentCancelling, AssignmentCancelled:
			return true
		default:
			return false
		}
	case AssignmentCancelling:
		switch to {
		case AssignmentCancelled, AssignmentFailed:
			return true
		default:
			return false
		}
	default:
		return false
	}
}
