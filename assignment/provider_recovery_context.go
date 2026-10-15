package assignment

import (
	"strconv"
	"strings"
)

// RecoveryContext is a bounded observation. Dispatched is not recovered, and
// recovery verification never implies semantic completion of the user's task.
type RecoveryContext struct {
	State            string `json:"state"`
	Reason           string `json:"reason"`
	Provider         string `json:"provider"`
	Circuit          string `json:"circuit"`
	OutcomeCertainty string `json:"outcome_certainty"`
	Attempt          int    `json:"attempt"`
	SwitchCount      int    `json:"switch_count"`
}

const (
	RecoveryStateKey            = "recovery_state"
	RecoveryReasonKey           = "recovery_reason"
	RecoveryProviderKey         = "recovery_provider"
	RecoveryCircuitKey          = "recovery_circuit"
	RecoveryOutcomeCertaintyKey = "recovery_outcome_certainty"
	RecoveryAttemptKey          = "recovery_attempt"
	RecoverySwitchCountKey      = "recovery_switch_count"
)

func RecoveryStates() []string {
	return []string{"observed", "checking", "backoff", "switching", "dispatched", "recovered", "user_action_required", "unavailable", "cancelled", "unknown"}
}

func RecoveryReasons() []string {
	return []string{"none", "transient_preflight", "provider_unavailable", "capability_pending", "slot_exhausted", "unknown_outcome", "session_pinned", "not_authorized", "deadline_exceeded", "budget_exhausted", "cancelled", "security", "unknown"}
}

func RecoveryCircuits() []string { return []string{"closed", "open", "half_open", "unknown"} }

func RecoveryOutcomeCertainties() []string {
	return []string{"not_started", "started", "confirmed", "unknown"}
}

func (c RecoveryContext) Sanitized() RecoveryContext {
	c.State = boundedExecutionValue(c.State, RecoveryStates(), "unknown")
	c.Reason = boundedExecutionValue(c.Reason, RecoveryReasons(), "unknown")
	if !RecoveryProviderValid(c.Provider) {
		c.Provider = "unknown"
	}
	c.Circuit = boundedExecutionValue(c.Circuit, RecoveryCircuits(), "unknown")
	c.OutcomeCertainty = boundedExecutionValue(c.OutcomeCertainty, RecoveryOutcomeCertainties(), "unknown")
	if c.Attempt < 0 || c.Attempt > 3 {
		c.Attempt = 0
	}
	if c.SwitchCount < 0 || c.SwitchCount > 1 {
		c.SwitchCount = 0
	}
	return c
}

func (c RecoveryContext) Metadata() map[string]string {
	c = c.Sanitized()
	return map[string]string{
		RecoveryStateKey: c.State, RecoveryReasonKey: c.Reason,
		RecoveryProviderKey: c.Provider, RecoveryCircuitKey: c.Circuit,
		RecoveryOutcomeCertaintyKey: c.OutcomeCertainty,
		RecoveryAttemptKey:          strconv.Itoa(c.Attempt),
		RecoverySwitchCountKey:      strconv.Itoa(c.SwitchCount),
	}
}

func RecoveryContextFromMetadata(metadata map[string]string) *RecoveryContext {
	if metadata[RecoveryStateKey] == "" && metadata[RecoveryReasonKey] == "" && metadata[RecoveryProviderKey] == "" && metadata[RecoveryCircuitKey] == "" && metadata[RecoveryOutcomeCertaintyKey] == "" && metadata[RecoveryAttemptKey] == "" && metadata[RecoverySwitchCountKey] == "" {
		return nil
	}
	attempt, _ := strconv.Atoi(metadata[RecoveryAttemptKey])
	switches, _ := strconv.Atoi(metadata[RecoverySwitchCountKey])
	c := (RecoveryContext{State: metadata[RecoveryStateKey], Reason: metadata[RecoveryReasonKey], Provider: metadata[RecoveryProviderKey], Circuit: metadata[RecoveryCircuitKey], OutcomeCertainty: metadata[RecoveryOutcomeCertaintyKey], Attempt: attempt, SwitchCount: switches}).Sanitized()
	return &c
}

// SanitizeRecoveryMetadata retains other event fields but replaces each recovery
// field with its bounded representation and drops unknown reserved keys.
// It never mutates the caller's map.
func SanitizeRecoveryMetadata(metadata map[string]string) map[string]string {
	if metadata == nil {
		return nil
	}
	out := make(map[string]string, len(metadata))
	for key, value := range metadata {
		if strings.HasPrefix(key, "recovery_") {
			continue
		}
		out[key] = value
	}
	if c := RecoveryContextFromMetadata(metadata); c != nil {
		for key, value := range c.Metadata() {
			out[key] = value
		}
	}
	return out
}
