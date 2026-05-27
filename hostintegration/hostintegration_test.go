package hostintegration

import "testing"

func TestDistributionChannelContract(t *testing.T) {
	valid := []DistributionChannel{
		DistributionChannelDeveloperID,
		DistributionChannelMacAppStore,
		DistributionChannelMSIXSideload,
		DistributionChannelMSIXStore,
		DistributionChannelDevLocal,
	}
	for _, channel := range valid {
		if !channel.Valid() {
			t.Fatalf("%q should be a valid distribution channel", channel)
		}
	}
	if DistributionChannel("unknown").Valid() {
		t.Fatal("unknown distribution channel should be invalid")
	}

	storeManaged := map[DistributionChannel]bool{
		DistributionChannelDeveloperID:  false,
		DistributionChannelMacAppStore:  true,
		DistributionChannelMSIXSideload: false,
		DistributionChannelMSIXStore:    true,
		DistributionChannelDevLocal:     false,
	}
	for channel, want := range storeManaged {
		if got := channel.StoreManaged(); got != want {
			t.Fatalf("%q StoreManaged() = %v, want %v", channel, got, want)
		}
	}
}

func TestProviderRoutingStatusContract(t *testing.T) {
	valid := []ProviderRoutingStatus{
		ProviderRoutingAvailable,
		ProviderRoutingLoginRequired,
		ProviderRoutingUnsupported,
		ProviderRoutingStoreBlocked,
	}
	for _, status := range valid {
		if !status.Valid() {
			t.Fatalf("%q should be a valid provider routing status", status)
		}
	}
	if ProviderRoutingStatus("unknown").Valid() {
		t.Fatal("unknown provider routing status should be invalid")
	}
}
