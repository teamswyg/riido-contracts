package hostintegration

func NonOwnedSurfaces() []string {
	return []string{
		"app data root selection",
		"local IPC",
		"workspace grants",
		"consent ledgers",
		"review/demo mode",
		"privacy metadata allowlist artifacts",
		"provider executable discovery",
		"provider login probes",
		"provider process execution",
		"control-plane HTTP handlers",
		"Terraform, AWS accounts, or deployment evidence",
	}
}
