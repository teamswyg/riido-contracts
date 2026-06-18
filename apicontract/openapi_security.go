package apicontract

func securityForAuth(auth Auth) []map[string][]string {
	switch auth.Scheme {
	case "apiKey":
		return []map[string][]string{{"riidoAIAgentToken": []string{}}}
	default:
		return []map[string][]string{{"bearerAuth": []string{}}}
	}
}

func securitySchemesForIR(ir IRDocument) map[string]OpenAPISecurityScheme {
	schemes := map[string]OpenAPISecurityScheme{}
	for _, op := range ir.Operations {
		switch op.Auth.Scheme {
		case "apiKey":
			schemes["riidoAIAgentToken"] = OpenAPISecurityScheme{Type: "apiKey", In: "header", Name: op.Auth.Header}
		default:
			schemes["bearerAuth"] = OpenAPISecurityScheme{Type: "http", Scheme: "bearer", BearerFormat: "opaque"}
		}
	}
	return schemes
}
