package capability

// ProtocolCriticalArgs returns the provider CLI flags that belong to the
// adapter protocol itself and therefore cannot be overridden through free-form
// custom args.
func ProtocolCriticalArgs(kind ProtocolKind) []string {
	args := protocolCriticalArgs[kind]
	out := make([]string, len(args))
	copy(out, args)
	return out
}

var protocolCriticalArgs = map[ProtocolKind][]string{
	ProtocolClaudeStreamJSON: {
		"-p",
		"--print",
		"--output-format",
		"--input-format",
		"--permission-mode",
		"--mcp-config",
		"--strict-mcp-config",
		"--verbose",
	},
	ProtocolCodexExecJSONL: {
		"--json",
	},
	ProtocolCodexAppServer: {
		"--listen",
	},
	ProtocolOpenClawAgentJSON: {
		"--local",
		"--json",
		"--session-id",
		"--message",
		"--model",
		"--system-prompt",
	},
	ProtocolCursorAgentStreamJSON: {
		"-p",
		"--output-format",
		"--yolo",
	},
}
