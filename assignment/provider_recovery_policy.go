package assignment

import (
	"errors"
	"strings"
	"unicode"
)

// ProviderRecoveryPolicy authorizes bounded execution-before-start recovery.
// Fallbacks are explicit provider/model pairs, never inferred from task text.
// It does not authorize replay after provider dispatch or new permissions.
type ProviderRecoveryPolicy struct {
	MaxAttempts       int                         `json:"max_attempts,omitempty"`
	MaxElapsedSeconds int                         `json:"max_elapsed_seconds,omitempty"`
	MaxSwitches       int                         `json:"max_switches,omitempty"`
	Fallbacks         []ProviderRecoveryCandidate `json:"fallbacks,omitempty"`
}

type ProviderRecoveryCandidate struct {
	Provider string `json:"provider"`
	ModelID  string `json:"model_id"`
}

// Normalized returns an independent finite policy. Nil enables same-provider
// preflight retries only. Cross-provider recovery requires explicit opt-in.
func (p *ProviderRecoveryPolicy) Normalized() (ProviderRecoveryPolicy, error) {
	out := ProviderRecoveryPolicy{MaxAttempts: 3, MaxElapsedSeconds: 30}
	if p != nil {
		out = *CloneProviderRecoveryPolicy(p)
		if out.MaxAttempts == 0 {
			out.MaxAttempts = 3
		}
		if out.MaxElapsedSeconds == 0 {
			out.MaxElapsedSeconds = 30
		}
	}
	if out.MaxAttempts < 1 || out.MaxAttempts > 3 || out.MaxElapsedSeconds < 1 || out.MaxElapsedSeconds > 60 || out.MaxSwitches < 0 || out.MaxSwitches > 1 || len(out.Fallbacks) > 3 {
		return ProviderRecoveryPolicy{}, errors.New("invalid provider recovery budget")
	}
	if len(out.Fallbacks) > 0 && (out.MaxSwitches != 1 || out.MaxAttempts < 2) || out.MaxSwitches == 1 && len(out.Fallbacks) == 0 {
		return ProviderRecoveryPolicy{}, errors.New("provider switch requires explicit candidates and budget")
	}
	seen := map[string]bool{}
	for _, c := range out.Fallbacks {
		if !RecoveryProviderValid(c.Provider) || c.ModelID == "" || len(c.ModelID) > 128 || strings.IndexFunc(c.ModelID, func(r rune) bool { return unicode.IsSpace(r) || unicode.IsControl(r) }) >= 0 || seen[c.Provider] {
			return ProviderRecoveryPolicy{}, errors.New("invalid provider recovery candidate")
		}
		seen[c.Provider] = true
	}
	return out, nil
}

func RecoveryProviderValid(provider string) bool {
	return provider == "claude" || provider == "codex" || provider == "cursor" || provider == "openclaw"
}

func CloneProviderRecoveryPolicy(p *ProviderRecoveryPolicy) *ProviderRecoveryPolicy {
	if p == nil {
		return nil
	}
	out := *p
	out.Fallbacks = append([]ProviderRecoveryCandidate(nil), p.Fallbacks...)
	return &out
}
