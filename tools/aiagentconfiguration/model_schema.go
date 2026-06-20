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

func findProperty(schema schema, name string) property {
	for _, prop := range schema.Properties {
		if prop.Name == name {
			return prop
		}
	}
	return property{}
}
