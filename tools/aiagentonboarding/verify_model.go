package main

func verifyModel(model model) error {
	if err := verifyCounts(model); err != nil {
		return err
	}
	if err := verifyOperations(model); err != nil {
		return err
	}
	if err := verifySchemas(model); err != nil {
		return err
	}
	if err := verifyFixtureRows(model.Manifest.FixtureRows); err != nil {
		return err
	}
	if err := verifyProjectionFlags(model); err != nil {
		return err
	}
	return nil
}
