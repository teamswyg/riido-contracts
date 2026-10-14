package main

import "fmt"

func verifyEnum(model model) error {
	exp := model.Manifest.RequiredEnum
	if model.Enum.Name != exp.Name {
		return fmt.Errorf("enum %s missing", exp.Name)
	}
	for _, value := range exp.Values {
		if !enumHasValue(model.Enum, value) {
			return fmt.Errorf("enum %s value %s missing", exp.Name, value)
		}
	}
	return nil
}
