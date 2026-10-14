package main

func addEnums(model model, dsl contractFixture) model {
	for _, exp := range model.Manifest.RequiredEnums {
		model.Enums = append(model.Enums, findEnum(dsl.Enums, exp.Name))
	}
	return model
}

func findEnum(enums []enumSpec, name string) enumSpec {
	for _, enum := range enums {
		if enum.Name == name {
			return enum
		}
	}
	return enumSpec{}
}

func enumValues(enum enumSpec) []string {
	values := make([]string, 0, len(enum.Values))
	for _, value := range enum.Values {
		values = append(values, value.Value)
	}
	return values
}
