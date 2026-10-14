package catalog

import "github.com/teamswyg/riido-contracts/provider/capability"

type Kind = capability.ProviderKind

const (
	KindClaude     Kind = "claude"
	KindClaudeCode Kind = "claude_code"
	KindCodex      Kind = "codex"
	KindCursor     Kind = "cursor"
	KindGemini     Kind = "gemini"
	KindOpenClaw   Kind = "openclaw"
)
