package main

import "fmt"

func verifyEnums(model model) error {
	for _, exp := range model.Manifest.RequiredEnums {
		enum := findEnum(model.Enums, exp.Name)
		if enum.Name == "" {
			return fmt.Errorf("enum %s missing", exp.Name)
		}
		if !sameStrings(enumValues(enum), exp.Values) {
			return fmt.Errorf("enum %s values mismatch", exp.Name)
		}
	}
	return nil
}

func sameStrings(left, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}
