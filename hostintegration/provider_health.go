package hostintegration

// ProviderHealthStatus describes observed provider readiness independently of
// the routing decision made by the control plane.
type ProviderHealthStatus string

const (
	ProviderHealthHealthy     ProviderHealthStatus = "healthy"
	ProviderHealthDegraded    ProviderHealthStatus = "degraded"
	ProviderHealthUnavailable ProviderHealthStatus = "unavailable"
	ProviderHealthUnknown     ProviderHealthStatus = "unknown"
)

func (s ProviderHealthStatus) Valid() bool {
	for _, status := range ProviderHealthStatuses() {
		if s == status {
			return true
		}
	}
	return false
}

func ProviderHealthStatuses() []ProviderHealthStatus {
	return []ProviderHealthStatus{
		ProviderHealthHealthy,
		ProviderHealthDegraded,
		ProviderHealthUnavailable,
		ProviderHealthUnknown,
	}
}

// ProviderDiagnosticCode is a bounded, non-sensitive explanation for a
// non-healthy provider observation.
type ProviderDiagnosticCode string

const (
	ProviderDiagnosticNone                  ProviderDiagnosticCode = "none"
	ProviderDiagnosticExecutableMissing     ProviderDiagnosticCode = "executable-missing"
	ProviderDiagnosticLoginRequired         ProviderDiagnosticCode = "login-required"
	ProviderDiagnosticVersionUnsupported    ProviderDiagnosticCode = "version-unsupported"
	ProviderDiagnosticProbeFailed           ProviderDiagnosticCode = "probe-failed"
	ProviderDiagnosticAuthProbeFailed       ProviderDiagnosticCode = "auth-probe-failed"
	ProviderDiagnosticVersionProbeFailed    ProviderDiagnosticCode = "version-probe-failed"
	ProviderDiagnosticCapabilityProbeFailed ProviderDiagnosticCode = "capability-probe-failed"
	ProviderDiagnosticRuntimeError          ProviderDiagnosticCode = "runtime-error"
)

func (c ProviderDiagnosticCode) Valid() bool {
	for _, code := range ProviderDiagnosticCodes() {
		if c == code {
			return true
		}
	}
	return false
}

func ProviderDiagnosticCodes() []ProviderDiagnosticCode {
	return []ProviderDiagnosticCode{
		ProviderDiagnosticNone,
		ProviderDiagnosticExecutableMissing,
		ProviderDiagnosticLoginRequired,
		ProviderDiagnosticVersionUnsupported,
		ProviderDiagnosticProbeFailed,
		ProviderDiagnosticAuthProbeFailed,
		ProviderDiagnosticVersionProbeFailed,
		ProviderDiagnosticCapabilityProbeFailed,
		ProviderDiagnosticRuntimeError,
	}
}
