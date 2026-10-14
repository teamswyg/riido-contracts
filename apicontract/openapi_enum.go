package apicontract

func enumToOpenAPI(enum Enum) map[string]any {
	values := make([]string, 0, len(enum.Values))
	for _, value := range enum.Values {
		values = append(values, value.Value)
	}
	out := map[string]any{
		"type": enum.Type,
		"enum": values,
	}
	if enum.Description != "" {
		out["description"] = enum.Description
	}
	return out
}
