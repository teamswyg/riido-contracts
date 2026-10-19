package apicontract

import (
	"fmt"
	"regexp"
)

func validateStringConstraints(schema string, property Property) error {
	if property.MinLength == nil && property.Pattern == "" {
		return nil
	}
	if property.Type != "string" || property.Ref != "" {
		return fmt.Errorf("apicontract: schema %q property %q string constraints require string type", schema, property.Name)
	}
	if property.MinLength != nil && (*property.MinLength < 0 ||
		property.MaxLength != nil && *property.MinLength > *property.MaxLength) {
		return fmt.Errorf("apicontract: schema %q property %q min_length is outside its length bounds", schema, property.Name)
	}
	if property.Pattern != "" {
		if _, err := regexp.Compile(property.Pattern); err != nil {
			return fmt.Errorf("apicontract: schema %q property %q pattern is invalid", schema, property.Name)
		}
	}
	return nil
}
