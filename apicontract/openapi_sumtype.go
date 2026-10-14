package apicontract

func sumTypeToOpenAPI(sumType SumType) map[string]any {
	oneOf := make([]map[string]any, 0, len(sumType.Variants))
	mapping := map[string]string{}
	for _, variant := range sumType.Variants {
		oneOf = append(oneOf, refSchema(variant.Schema))
		mapping[variant.Kind] = "#/components/schemas/" + variant.Schema
	}
	out := map[string]any{
		"oneOf": oneOf,
		"discriminator": map[string]any{
			"propertyName": sumType.Discriminator,
			"mapping":      mapping,
		},
	}
	if sumType.Description != "" {
		out["description"] = sumType.Description
	}
	return out
}
