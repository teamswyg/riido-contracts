package assignment

// RecoveryCode classifies daemon-authored recovery handling in assignment event
// metadata. It is stored as a string to preserve existing event-history shape.
type RecoveryCode string

const (
	RecoveryFreshStartRefused RecoveryCode = "fresh_start_refused"
)

func (code RecoveryCode) String() string {
	return string(code)
}
