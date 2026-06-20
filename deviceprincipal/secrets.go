package deviceprincipal

func SecretNonExposureSinks() []string {
	return []string{
		"browser local storage",
		"cookies",
		"webview JavaScript state",
		"renderer IPC payloads",
		"command-line arguments",
		"daemon status output",
		"logs, task evidence, SSE events, or generated client responses",
	}
}
