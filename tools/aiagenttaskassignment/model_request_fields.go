package main

func requestFieldsAbsent(model model, dsl contractFixture) bool {
	for _, exp := range model.Manifest.RequiredOperations {
		if exp.RequestRef == "" {
			continue
		}
		schema := findSchema(dsl.Schemas, exp.RequestRef)
		if schema.Name == "" || hasForbiddenProperty(schema, model.Manifest.ForbiddenRequestFields) {
			return false
		}
	}
	return true
}

func hasForbiddenProperty(schema schema, forbidden []string) bool {
	for _, prop := range schema.Properties {
		for _, field := range forbidden {
			if prop.Name == field {
				return true
			}
		}
	}
	return false
}
