package main

func verifyModel(model model) error {
	for _, check := range modelChecks(model) {
		if err := check(); err != nil {
			return err
		}
	}
	return nil
}

func modelChecks(model model) []func() error {
	return []func() error{
		func() error { return verifyCounts(model) },
		func() error { return verifyBooleans(model) },
	}
}
