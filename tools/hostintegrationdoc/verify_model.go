package main

func verifyModel(model model) error {
	if err := verifyCounts(model); err != nil {
		return err
	}
	if !model.DistributionValid {
		return errInvalidDistribution
	}
	if !model.ProviderRoutingValid {
		return errInvalidProviderRouting
	}
	if !model.StoreManagedExclusive {
		return errInvalidStoreManaged
	}
	return nil
}
