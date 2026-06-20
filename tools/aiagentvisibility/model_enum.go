package main

func findEnum(enums []enumSpec, name string) enumSpec {
	for _, enum := range enums {
		if enum.Name == name {
			return enum
		}
	}
	return enumSpec{}
}

func enumHasValue(enum enumSpec, value string) bool {
	for _, got := range enum.Values {
		if got.Value == value {
			return true
		}
	}
	return false
}
