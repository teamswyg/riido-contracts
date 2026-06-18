package apicontract

import (
	"errors"
	"fmt"
	"strings"
)

func validateDSLEnums(enums []Enum, components map[string]struct{}) error {
	for _, enum := range enums {
		if strings.TrimSpace(enum.Name) == "" {
			return errors.New("apicontract: enum name is required")
		}
		if enum.Type != "string" {
			return fmt.Errorf("apicontract: enum %q must use string type", enum.Name)
		}
		if len(enum.Values) == 0 {
			return fmt.Errorf("apicontract: enum %q must define values", enum.Name)
		}
		if _, exists := components[enum.Name]; exists {
			return fmt.Errorf("apicontract: duplicate component %q", enum.Name)
		}
		components[enum.Name] = struct{}{}
		if err := validateDSLEnumValues(enum); err != nil {
			return err
		}
	}
	return nil
}

func validateDSLEnumValues(enum Enum) error {
	values := map[string]struct{}{}
	for _, value := range enum.Values {
		if strings.TrimSpace(value.Value) == "" {
			return fmt.Errorf("apicontract: enum %q has blank value", enum.Name)
		}
		if _, exists := values[value.Value]; exists {
			return fmt.Errorf("apicontract: enum %q has duplicate value %q", enum.Name, value.Value)
		}
		values[value.Value] = struct{}{}
	}
	return nil
}
