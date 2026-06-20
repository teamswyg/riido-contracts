package deviceprincipal

const (
	HeaderDeviceID     = "X-Riido-Device-ID"
	HeaderDeviceSecret = "X-Riido-Device-Secret"
	HeaderAIAgentToken = "X-Riido-AI-Agent-Token"
)

func DaemonCredentialHeaders() []string {
	return []string{HeaderDeviceID, HeaderDeviceSecret}
}

func ClientCredentialHeaders() []string {
	return []string{HeaderAIAgentToken}
}
