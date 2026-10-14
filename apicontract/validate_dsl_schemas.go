package apicontract

import (
	"errors"
	"fmt"
	"strings"
)

func validateDSLSchemas(schemas []Schema, index dslValidationIndex) error {
	for _, schema := range schemas {
		if strings.TrimSpace(schema.Name) == "" {
			return errors.New("apicontract: schema name is required")
		}
		if _, exists := index.components[schema.Name]; exists {
			return fmt.Errorf("apicontract: duplicate component %q", schema.Name)
		}
		index.schemas[schema.Name] = struct{}{}
		index.components[schema.Name] = struct{}{}
		if schema.Type != "object" {
			return fmt.Errorf("apicontract: schema %q must be object", schema.Name)
		}
		if err := validateDSLSchemaProperties(schema, index.components); err != nil {
			return err
		}
	}
	return nil
}

func validateDSLSchemaProperties(schema Schema, components map[string]struct{}) error {
	for _, property := range schema.Properties {
		if strings.TrimSpace(property.Name) == "" {
			return fmt.Errorf("apicontract: schema %q has blank property name", schema.Name)
		}
		if property.MaxLength != nil && *property.MaxLength <= 0 {
			return fmt.Errorf("apicontract: schema %q property %q max_length must be positive", schema.Name, property.Name)
		}
		if err := validatePropertyRef(schema.Name, property, components); err != nil {
			return err
		}
	}
	return nil
}
