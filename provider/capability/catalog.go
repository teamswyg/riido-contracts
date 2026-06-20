package capability

var protocolKinds = []ProtocolKind{
	ProtocolClaudeStreamJSON,
	ProtocolCodexExecJSONL,
	ProtocolCodexAppServer,
	ProtocolClaudeCompatibleWrapper,
	ProtocolOpenClawAgentJSON,
	ProtocolCursorAgentStreamJSON,
}

var eventStreamFormats = []EventStreamFormat{
	EventStreamFormatUnknown,
	EventStreamFormatTextOnly,
	EventStreamFormatNDJSON,
	EventStreamFormatJSONRPCNotifications,
}

var protocolMaturities = []ProtocolMaturity{
	ProtocolMaturityUnknown,
	ProtocolMaturityStable,
	ProtocolMaturityExperimental,
	ProtocolMaturityDeprecated,
}

var compatibilityStatuses = []CompatibilityStatus{
	CompatSupported,
	CompatDegraded,
	CompatExperimental,
	CompatBlocked,
}

func AllProtocolKinds() []ProtocolKind {
	return append([]ProtocolKind(nil), protocolKinds...)
}

func AllEventStreamFormats() []EventStreamFormat {
	return append([]EventStreamFormat(nil), eventStreamFormats...)
}

func AllProtocolMaturities() []ProtocolMaturity {
	return append([]ProtocolMaturity(nil), protocolMaturities...)
}

func AllCompatibilityStatuses() []CompatibilityStatus {
	return append([]CompatibilityStatus(nil), compatibilityStatuses...)
}
