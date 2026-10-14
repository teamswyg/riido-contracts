package main

import "github.com/teamswyg/riido-contracts/ir"

func eventScopes() []ir.EventScope {
	return []ir.EventScope{ir.EventScopeSystem, ir.EventScopeRuntime, ir.EventScopeTask, ir.EventScopeRun}
}

func commonRequiredFields() []string {
	return []string{"EventID", "OccurredAt", "EventSchemaVersion", "Type", "ActorKind", "RiidoDaemonVersion", "PolicyBundleVersion"}
}

func runRequiredFields() []string {
	return []string{"TaskID", "RunID", "RuntimeID", "CapabilityFingerprint", "ProviderKind", "ProtocolKind", "ProviderVersion", "AdapterID", "AdapterVersion", "ProtocolVersion"}
}

func fakePlaceholderFields() []string {
	return []string{"RuntimeID", "CapabilityFingerprint", "NativeConfigVersion", "ProviderVersion", "AdapterID", "AdapterVersion", "ProtocolVersion", "TaskID", "RunID"}
}
