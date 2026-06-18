package ir

func fieldValue(e CanonicalEvent, name string) string {
	switch name {
	case "TaskID":
		return e.TaskID
	case "RunID":
		return e.RunID
	case "RuntimeID":
		return e.RuntimeID
	case "CapabilityFingerprint":
		return e.CapabilityFingerprint
	case "ProviderKind":
		return e.ProviderKind
	case "ProtocolKind":
		return e.ProtocolKind
	case "ProviderVersion":
		return e.ProviderVersion
	case "AdapterID":
		return e.AdapterID
	case "AdapterVersion":
		return e.AdapterVersion
	case "ProtocolVersion":
		return e.ProtocolVersion
	case "NativeConfigVersion":
		return e.NativeConfigVersion
	}
	return ""
}
