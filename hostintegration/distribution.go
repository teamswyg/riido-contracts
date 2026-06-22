package hostintegration

import "slices"

// DistributionChannel is the package artifact identity that constrains which
// host surfaces may be used.
type DistributionChannel string

const (
	DistributionChannelDeveloperID  DistributionChannel = "developer-id"
	DistributionChannelMacAppStore  DistributionChannel = "mac-app-store"
	DistributionChannelMSIXSideload DistributionChannel = "msix-sideload"
	DistributionChannelMSIXStore    DistributionChannel = "msix-store"
	DistributionChannelDevLocal     DistributionChannel = "dev-local"
)

// Valid reports whether channel is one of the SSOT-defined distribution
// channels.
func (c DistributionChannel) Valid() bool {
	return slices.Contains(DistributionChannels(), c)
}

// StoreManaged reports whether the channel is subject to app store review
// constraints.
func (c DistributionChannel) StoreManaged() bool {
	return slices.Contains(StoreManagedDistributionChannels(), c)
}

func DistributionChannels() []DistributionChannel {
	return []DistributionChannel{
		DistributionChannelDeveloperID,
		DistributionChannelMacAppStore,
		DistributionChannelMSIXSideload,
		DistributionChannelMSIXStore,
		DistributionChannelDevLocal,
	}
}

func StoreManagedDistributionChannels() []DistributionChannel {
	return []DistributionChannel{
		DistributionChannelMacAppStore,
		DistributionChannelMSIXStore,
	}
}
