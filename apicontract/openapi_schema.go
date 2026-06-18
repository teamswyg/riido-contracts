package apicontract

func schemaToOpenAPI(schema Schema) map[string]any {
	out := map[string]any{"type": schema.Type}
	if schema.Description != "" {
		out["description"] = schema.Description
	}
	if len(schema.Required) > 0 {
		out["required"] = append([]string(nil), schema.Required...)
	}
	if len(schema.Properties) > 0 {
		properties := map[string]any{}
		for _, property := range schema.Properties {
			properties[property.Name] = propertyToOpenAPI(property)
		}
		out["properties"] = properties
	}
	return out
}
