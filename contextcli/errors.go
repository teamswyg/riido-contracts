package contextcli

// ExitCode is the process exit status of a work-context command.
type ExitCode int

const (
	ExitOK           ExitCode = 0
	ExitUsage        ExitCode = 2
	ExitAuthRequired ExitCode = 3
	ExitPermission   ExitCode = 4
	ExitNotFound     ExitCode = 5
	ExitUnavailable  ExitCode = 6
	ExitRateLimited  ExitCode = 7
	ExitUnsupported  ExitCode = 8
	ExitIncomplete   ExitCode = 9
	ExitConflict     ExitCode = 10
)

// ErrorCode is the stable machine-readable failure code in the envelope.
type ErrorCode string

const (
	CodeUsage                   ErrorCode = "USAGE"
	CodeSecretInArgv            ErrorCode = "SECRET_IN_ARGV"
	CodeAuthRequired            ErrorCode = "AUTH_REQUIRED"
	CodePermissionDenied        ErrorCode = "PERMISSION_DENIED"
	CodeOutOfRunScope           ErrorCode = "OUT_OF_RUN_SCOPE"
	CodeNotFound                ErrorCode = "NOT_FOUND"
	CodeSnapshotExpired         ErrorCode = "SNAPSHOT_EXPIRED"
	CodeUnavailable             ErrorCode = "UNAVAILABLE"
	CodeDaemonUnreachable       ErrorCode = "DAEMON_UNREACHABLE"
	CodeRateLimited             ErrorCode = "RATE_LIMITED"
	CodeUnsupported             ErrorCode = "UNSUPPORTED"
	CodeContextReadUnsupported  ErrorCode = "CONTEXT_READ_UNSUPPORTED"
	CodeUserModeUnavailable     ErrorCode = "USER_MODE_UNAVAILABLE"
	CodeControlPlaneUnsupported ErrorCode = "CONTROL_PLANE_UNSUPPORTED"
	CodeIncomplete              ErrorCode = "INCOMPLETE"
	CodeConflict                ErrorCode = "CONFLICT"
	CodeLeaseLost               ErrorCode = "LEASE_LOST"
)
