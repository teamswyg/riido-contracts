package assignment

// ConfirmationReason describes human intervention within an active execution.
type ConfirmationReason string

const (
	ConfirmationAnswerRequired    ConfirmationReason = "answer_required"
	ConfirmationApprovalRequired  ConfirmationReason = "approval_required"
	EventAssignmentWaitingForUser                    = "assignment_waiting_for_user"
)

func (reason ConfirmationReason) Valid() bool {
	return reason == ConfirmationAnswerRequired || reason == ConfirmationApprovalRequired
}
