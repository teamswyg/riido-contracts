package deviceprincipal

const (
	RuntimeSnapshotIntervalSeconds = 5
	RuntimeStaleAfterSeconds       = 20
)

func DependencyPhrases() []string {
	return []string{
		"daemon runtime snapshot every 5 seconds",
		"daemon process facts",
		"not been refreshed for 20 seconds",
		"stale runtimes must not be used for newly derived `AgentRuntimeBinding`",
	}
}
