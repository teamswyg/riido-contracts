package main

func selectSchemas(schemas []schema, exps []schemaExpectation) []schema {
	out := make([]schema, 0, len(exps))
	for _, exp := range exps {
		out = append(out, findSchema(schemas, exp.Schema))
	}
	return out
}

func findSchema(schemas []schema, name string) schema {
	for _, schema := range schemas {
		if schema.Name == name {
			return schema
		}
	}
	return schema{}
}

func schemaHasField(schema schema, name string) bool {
	for _, prop := range schema.Properties {
		if prop.Name == name {
			return true
		}
	}
	return false
}
