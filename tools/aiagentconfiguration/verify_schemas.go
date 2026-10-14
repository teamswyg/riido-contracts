package main

import "fmt"

func verifySchemas(model model) error {
	for _, exp := range model.Manifest.RequiredSchemaFields {
		schema := findSchema(model.Schemas, exp.Schema)
		if schema.Name == "" {
			return fmt.Errorf("schema %s missing", exp.Schema)
		}
		for _, field := range exp.Fields {
			if findProperty(schema, field).Name == "" {
				return fmt.Errorf("schema %s missing field %s", exp.Schema, field)
			}
		}
	}
	return verifyBoundedFields(model)
}

func verifyBoundedFields(model model) error {
	for _, bounded := range model.Manifest.BoundedFields {
		schema := findSchema(model.Schemas, bounded.Schema)
		prop := findProperty(schema, bounded.Field)
		if prop.MaxLength != bounded.MaxLength {
			return fmt.Errorf("%s.%s max = %d", bounded.Schema, bounded.Field, prop.MaxLength)
		}
	}
	return nil
}
