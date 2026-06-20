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
	if err := verifyPolicies(model); err != nil {
		return err
	}
	if err := verifyEnums(model); err != nil {
		return err
	}
	return verifyProjection(model)
}
