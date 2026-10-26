package contextcli

// ExitCode maps an error code to the documented process exit status.
func (c ErrorCode) ExitCode() ExitCode {
	switch c {
	case CodeUsage, CodeSecretInArgv:
		return ExitUsage
	case CodeAuthRequired:
		return ExitAuthRequired
	case CodePermissionDenied, CodeOutOfRunScope:
		return ExitPermission
	case CodeNotFound, CodeSnapshotExpired:
		return ExitNotFound
	case CodeUnavailable, CodeDaemonUnreachable:
		return ExitUnavailable
	case CodeRateLimited:
		return ExitRateLimited
	case CodeUnsupported, CodeContextReadUnsupported, CodeUserModeUnavailable, CodeControlPlaneUnsupported:
		return ExitUnsupported
	case CodeIncomplete:
		return ExitIncomplete
	case CodeConflict, CodeLeaseLost:
		return ExitConflict
	default:
		return ExitUnsupported
	}
}
