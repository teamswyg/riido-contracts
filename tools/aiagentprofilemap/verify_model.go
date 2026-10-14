package main

func verifyModel(model model) error {
	if err := verifyCounts(model); err != nil {
		return err
	}
	if err := verifyOperation(model); err != nil {
		return err
	}
	if err := verifyPolicy(model); err != nil {
		return err
	}
	if err := verifySchemas(model); err != nil {
		return err
	}
	if err := verifyProjection(model); err != nil {
		return err
	}
	return verifyMapShape(model)
}
