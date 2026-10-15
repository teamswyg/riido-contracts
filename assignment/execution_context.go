package assignment

// ExecutionContext is the bounded, non-sensitive view of a host execution.
// Absolute directories, verification commands and provider output stay local.
// Empty/absent context means legacy data, never verified execution.
type ExecutionContext struct {
	WorkspaceKind      string `json:"workspace_kind"`
	VerificationStatus string `json:"verification_status"`
	DiagnosticCode     string `json:"diagnostic_code"`
	NextAction         string `json:"next_action"`
}

const (
	ExecutionWorkspaceKindKey = "execution_workspace_kind"
	ExecutionVerificationKey  = "execution_verification_status"
	ExecutionDiagnosticKey    = "execution_diagnostic_code"
	ExecutionNextActionKey    = "execution_next_action"
)

func ExecutionWorkspaceKinds() []string {
	return []string{"unknown", "isolated", "repository_clone", "assignment_worktree", "native_resume"}
}

func ExecutionVerificationStatuses() []string {
	return []string{"not_configured", "pending", "passed", "failed", "cancelled", "unknown"}
}

func ExecutionDiagnosticCodes() []string {
	return []string{"none", "user_cancelled", "run_timeout", "workspace_preparation_failed", "session_scope_conflict", "host_verification_failed", "provider_blocked", "provider_failed", "process_aborted", "unknown"}
}

func ExecutionNextActions() []string {
	return []string{"none", "check_workspace", "review_session", "review_result", "check_provider", "review_permissions"}
}

func boundedExecutionValue(value string, allowed []string, fallback string) string {
	for _, candidate := range allowed {
		if value == candidate {
			return value
		}
	}
	return fallback
}

// Sanitized rejects free-form values at the shared diagnostic boundary.
func (c ExecutionContext) Sanitized() ExecutionContext {
	c.WorkspaceKind = boundedExecutionValue(c.WorkspaceKind, ExecutionWorkspaceKinds(), "unknown")
	c.VerificationStatus = boundedExecutionValue(c.VerificationStatus, ExecutionVerificationStatuses(), "unknown")
	c.DiagnosticCode = boundedExecutionValue(c.DiagnosticCode, ExecutionDiagnosticCodes(), "unknown")
	c.NextAction = boundedExecutionValue(c.NextAction, ExecutionNextActions(), "review_result")
	return c
}

func (c ExecutionContext) Metadata() map[string]string {
	c = c.Sanitized()
	return map[string]string{
		ExecutionWorkspaceKindKey: c.WorkspaceKind,
		ExecutionVerificationKey:  c.VerificationStatus,
		ExecutionDiagnosticKey:    c.DiagnosticCode,
		ExecutionNextActionKey:    c.NextAction,
	}
}

func ExecutionContextFromMetadata(metadata map[string]string) *ExecutionContext {
	if metadata[ExecutionWorkspaceKindKey] == "" && metadata[ExecutionVerificationKey] == "" && metadata[ExecutionDiagnosticKey] == "" && metadata[ExecutionNextActionKey] == "" {
		return nil
	}
	c := ExecutionContext{
		WorkspaceKind:      metadata[ExecutionWorkspaceKindKey],
		VerificationStatus: metadata[ExecutionVerificationKey],
		DiagnosticCode:     metadata[ExecutionDiagnosticKey],
		NextAction:         metadata[ExecutionNextActionKey],
	}.Sanitized()
	return &c
}
