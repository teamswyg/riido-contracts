package main

import "fmt"

func verifySchemas(model model) error {
	for _, exp := range model.Manifest.RequiredSchemaFields {
		schema := findSchema(model.Schemas, exp.Schema)
		if schema.Name == "" {
			return fmt.Errorf("schema %s missing", exp.Schema)
		}
		if err := verifySchemaFields(schema, exp.Fields); err != nil {
			return err
		}
	}
	return nil
}

func verifySchemaFields(schema schema, fields []string) error {
	for _, field := range fields {
		if !schemaHasField(schema, field) {
			return fmt.Errorf("schema %s field %s missing", schema.Name, field)
		}
		if schemaRequires(schema, field) && !schemaHasField(schema, field) {
			return fmt.Errorf("schema %s required field %s has no property", schema.Name, field)
		}
	}
	return nil
}
