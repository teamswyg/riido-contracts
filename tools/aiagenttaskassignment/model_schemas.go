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
	for _, field := range schema.Properties {
		if field.Name == name {
			return true
		}
	}
	return false
}

func schemaRequires(schema schema, name string) bool {
	for _, required := range schema.Required {
		if required == name {
			return true
		}
	}
	return false
}
