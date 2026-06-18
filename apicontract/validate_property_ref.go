package apicontract

import (
	"fmt"
)

func validatePropertyRef(schemaName string, property Property, components map[string]struct{}) error {
	if property.Ref != "" {
		if _, ok := components[property.Ref]; !ok {
			return fmt.Errorf("apicontract: schema %q property %q ref %q is missing", schemaName, property.Name, property.Ref)
		}
	}
	if property.AdditionalPropertiesRef != "" {
		if _, ok := components[property.AdditionalPropertiesRef]; !ok {
			return fmt.Errorf("apicontract: schema %q property %q additional_properties_ref %q is missing", schemaName, property.Name, property.AdditionalPropertiesRef)
		}
	}
	if property.Items != nil {
		return validatePropertyRef(schemaName, *property.Items, components)
	}
	return nil
}
