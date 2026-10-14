package main

import "fmt"

func verifySchemas(model model) error {
	for _, exp := range model.Manifest.RequiredSchemaFields {
		schema := findSchema(model.Schemas, exp.Schema)
		if schema.Name == "" {
			return fmt.Errorf("schema %s missing", exp.Schema)
		}
		for _, field := range exp.Fields {
			if !schemaHasField(schema, field) {
				return fmt.Errorf("schema %s field %s missing", schema.Name, field)
			}
		}
	}
	return nil
}
