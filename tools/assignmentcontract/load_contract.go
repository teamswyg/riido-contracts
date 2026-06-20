package main

import "path/filepath"

func loadContract(path string) (contract, error) {
	c, err := loadJSON[contract](path, "assignment contract")
	if err != nil {
		return contract{}, err
	}
	if err := loadContractIncludes(filepath.Dir(path), &c); err != nil {
		return contract{}, err
	}
	return c, nil
}

func loadContractIncludes(base string, c *contract) error {
	if err := loadStateIncludes(base, c); err != nil {
		return err
	}
	if err := loadNamedValueIncludes(base, c); err != nil {
		return err
	}
	if err := loadSingleIncludes(base, c); err != nil {
		return err
	}
	return loadPayloadIncludes(base, c)
}
