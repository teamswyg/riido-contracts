package main

func verifyModel(model model) error {
	if err := verifyCounts(model); err != nil {
		return err
	}
	if err := verifyPolicyPrefixes(model); err != nil {
		return err
	}
	if err := verifyBindingFields(model); err != nil {
		return err
	}
	if err := verifyHeaderBoundary(model); err != nil {
		return err
	}
	return nil
}
