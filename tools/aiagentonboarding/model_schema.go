package main

func findSchema(schemas []schema, name string) schema {
	for _, schema := range schemas {
		if schema.Name == name {
			return schema
		}
	}
	return schema{}
}

func propertyNames(schema schema) []string {
	names := make([]string, 0, len(schema.Properties))
	for _, prop := range schema.Properties {
		names = append(names, prop.Name)
	}
	return names
}

func hasProperty(schema schema, name string) bool {
	for _, prop := range schema.Properties {
		if prop.Name == name {
			return true
		}
	}
	return false
}
