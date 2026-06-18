package apicontract

func propertyToOpenAPI(property Property) map[string]any {
	if property.Ref != "" {
		out := refSchema(property.Ref)
		if property.Description != "" {
			out["description"] = property.Description
		}
		return out
	}
	out := openAPIPropertyScalar(property)
	if property.Items != nil {
		out["items"] = propertyToOpenAPI(*property.Items)
	}
	if property.AdditionalProperties {
		out["additionalProperties"] = true
	}
	if property.AdditionalPropertiesRef != "" {
		out["additionalProperties"] = refSchema(property.AdditionalPropertiesRef)
	}
	return out
}

func openAPIPropertyScalar(property Property) map[string]any {
	out := map[string]any{}
	if property.Type != "" {
		out["type"] = property.Type
	}
	if property.Description != "" {
		out["description"] = property.Description
	}
	if property.Format != "" {
		out["format"] = property.Format
	}
	if property.MaxLength != nil {
		out["maxLength"] = *property.MaxLength
	}
	if len(property.Enum) > 0 {
		out["enum"] = append([]string(nil), property.Enum...)
	}
	return out
}
