package apicontract

func responseForSchema(name, description, contentType string) OpenAPIResponse {
	if contentType == "" {
		contentType = "application/json"
	}
	return OpenAPIResponse{
		Description: description,
		Content: map[string]OpenAPIMedia{
			contentType: {Schema: refSchema(name)},
		},
	}
}

func statusDescription(status int) string {
	switch status {
	case 200:
		return "OK"
	case 201:
		return "Created"
	case 202:
		return "Accepted"
	default:
		return "response"
	}
}

func jsonContent(name string) map[string]OpenAPIMedia {
	return map[string]OpenAPIMedia{
		"application/json": {Schema: refSchema(name)},
	}
}
