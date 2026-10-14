package assignment

type ApprovalStatus string

type ApprovalDecision string

const (
	ApprovalContractOwner         = "contracts-control-plane-client"
	ApprovalTimeoutTerminalStatus = ApprovalTimedOut
)

func (value ApprovalStatus) IsTerminal() bool {
	switch value.Code() {
	case ApprovalStatusCodeApproved, ApprovalStatusCodeDenied, ApprovalStatusCodeTimedOut:
		return true
	default:
		return false
	}
}
