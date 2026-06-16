package catalog

import (
	"strings"

	"github.com/teamswyg/riido-contracts/provider/capability"
)

type Kind = capability.ProviderKind

const (
	KindClaude     Kind = "claude"
	KindClaudeCode Kind = "claude_code"
	KindCodex      Kind = "codex"
	KindCursor     Kind = "cursor"
	KindGemini     Kind = "gemini"
	KindOpenClaw   Kind = "openclaw"
)

const (
	DefaultClaudeModelID   = "claude-default"
	DefaultCodexModelID    = "codex-default"
	DefaultCursorModelID   = "cursor-auto"
	DefaultGeminiModelID   = "gemini-default"
	DefaultOpenClawModelID = "openclaw-default"
	DefaultRuntimeModelID  = "runtime-default"
)

func Normalize(value string) Kind {
	normalized := strings.TrimSpace(strings.ToLower(value))
	switch normalized {
	case "claude-code":
		return KindClaudeCode
	default:
		return Kind(normalized)
	}
}

func String(value string) string {
	return string(Normalize(value))
}

func IsCodex(value string) bool {
	return Normalize(value) == KindCodex
}

func IsClaudeFamily(value string) bool {
	switch Normalize(value) {
	case KindClaude, KindClaudeCode:
		return true
	default:
		return false
	}
}

func DefaultModelID(value string) string {
	switch Normalize(value) {
	case KindClaude, KindClaudeCode:
		return DefaultClaudeModelID
	case KindCodex:
		return DefaultCodexModelID
	case KindCursor:
		return DefaultCursorModelID
	case KindGemini:
		return DefaultGeminiModelID
	case KindOpenClaw:
		return DefaultOpenClawModelID
	default:
		return DefaultRuntimeModelID
	}
}

func ModelOverride(provider, modelID string) string {
	modelID = strings.TrimSpace(modelID)
	if modelID == "" || modelID == DefaultModelID(provider) {
		return ""
	}
	return modelID
}
