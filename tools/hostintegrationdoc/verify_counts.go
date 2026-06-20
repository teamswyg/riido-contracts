package main

import "fmt"

func verifyCounts(model model) error {
	m := model.Manifest
	checks := map[string][2]int{
		"distribution channels": {len(model.DistributionChannels), m.ExpectedDistributionChannels},
		"store managed":         {len(model.StoreManagedChannels), m.ExpectedStoreManagedChannels},
		"provider statuses":     {len(model.ProviderStatuses), m.ExpectedProviderRoutingStatus},
		"non-owned surfaces":    {len(model.NonOwnedSurfaces), m.ExpectedNonOwnedSurfaces},
	}
	for name, values := range checks {
		if values[0] != values[1] {
			return fmt.Errorf("%s = %d, want %d", name, values[0], values[1])
		}
	}
	return nil
}
