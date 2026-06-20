package main

func addSchemas(model model, dsl contractFixture) model {
	for _, exp := range model.Manifest.RequiredSchemaFields {
		model.Schemas = append(model.Schemas, findSchema(dsl.Schemas, exp.Schema))
	}
	return model
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
