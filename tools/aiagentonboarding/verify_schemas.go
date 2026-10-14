package main

import "fmt"

func verifySchemas(model model) error {
	if model.FixtureSchema.Name == "" || model.ListSchema.Name == "" {
		return fmt.Errorf("onboarding fixture schemas are required")
	}
	if model.CreateRequestSchema.Name == "" {
		return fmt.Errorf("create request schema is required")
	}
	if err := verifySchemaFields(model.FixtureSchema, model.Manifest.RequiredFixtureFields); err != nil {
		return err
	}
	return verifySchemaFields(model.CreateRequestSchema, model.Manifest.RequiredCreateRequestFields)
}

func verifySchemaFields(schema schema, fields []string) error {
	for _, field := range fields {
		if !hasProperty(schema, field) {
			return fmt.Errorf("schema %s missing field %s", schema.Name, field)
		}
	}
	return nil
}
