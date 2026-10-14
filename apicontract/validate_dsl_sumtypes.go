package apicontract

import (
	"errors"
	"fmt"
	"strings"
)

func validateDSLSumTypes(sumTypes []SumType, components map[string]struct{}) error {
	for _, sumType := range sumTypes {
		if strings.TrimSpace(sumType.Name) == "" {
			return errors.New("apicontract: sum_type name is required")
		}
		if strings.TrimSpace(sumType.Discriminator) == "" {
			return fmt.Errorf("apicontract: sum_type %q discriminator is required", sumType.Name)
		}
		if len(sumType.Variants) == 0 {
			return fmt.Errorf("apicontract: sum_type %q must define variants", sumType.Name)
		}
		if _, exists := components[sumType.Name]; exists {
			return fmt.Errorf("apicontract: duplicate component %q", sumType.Name)
		}
		components[sumType.Name] = struct{}{}
		if err := validateDSLSumTypeVariants(sumType); err != nil {
			return err
		}
	}
	return nil
}

func validateDSLSumTypeVariants(sumType SumType) error {
	kinds := map[string]struct{}{}
	for _, variant := range sumType.Variants {
		if strings.TrimSpace(variant.Kind) == "" || strings.TrimSpace(variant.Schema) == "" {
			return fmt.Errorf("apicontract: sum_type %q variant kind and schema are required", sumType.Name)
		}
		if _, exists := kinds[variant.Kind]; exists {
			return fmt.Errorf("apicontract: sum_type %q has duplicate variant %q", sumType.Name, variant.Kind)
		}
		kinds[variant.Kind] = struct{}{}
	}
	return nil
}

func validateDSLSumTypeVariantSchemas(sumTypes []SumType, schemas map[string]struct{}) error {
	for _, sumType := range sumTypes {
		for _, variant := range sumType.Variants {
			if _, ok := schemas[variant.Schema]; !ok {
				return fmt.Errorf("apicontract: sum_type %q variant schema %q is missing", sumType.Name, variant.Schema)
			}
		}
	}
	return nil
}
