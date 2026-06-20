package hostintegration

// ProviderRoutingStatus is the server-facing provider status vocabulary shared
// by C11 metadata collection and C10 routing/sync API contracts.
type ProviderRoutingStatus string

const (
	ProviderRoutingAvailable     ProviderRoutingStatus = "available"
	ProviderRoutingLoginRequired ProviderRoutingStatus = "login-required"
	ProviderRoutingUnsupported   ProviderRoutingStatus = "unsupported"
	ProviderRoutingStoreBlocked  ProviderRoutingStatus = "store-blocked"
)

func (s ProviderRoutingStatus) Valid() bool {
	for _, status := range ProviderRoutingStatuses() {
		if s == status {
			return true
		}
	}
	return false
}

func ProviderRoutingStatuses() []ProviderRoutingStatus {
	return []ProviderRoutingStatus{
		ProviderRoutingAvailable,
		ProviderRoutingLoginRequired,
		ProviderRoutingUnsupported,
		ProviderRoutingStoreBlocked,
	}
}
