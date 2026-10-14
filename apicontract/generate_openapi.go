package apicontract

import (
	"strconv"
	"strings"
)

func GenerateOpenAPI(ir IRDocument) (OpenAPISpec, error) {
	if err := validateIR(ir); err != nil {
		return OpenAPISpec{}, err
	}
	components := map[string]map[string]any{}
	for _, enum := range ir.Enums {
		components[enum.Name] = enumToOpenAPI(enum)
	}
	for _, sumType := range ir.SumTypes {
		components[sumType.Name] = sumTypeToOpenAPI(sumType)
	}
	for _, component := range ir.Components {
		components[component.Name] = schemaToOpenAPI(component.Schema)
	}
	paths := map[string]OpenAPIPath{}
	for _, op := range ir.Operations {
		method := strings.ToLower(op.Method)
		if paths[op.Path] == nil {
			paths[op.Path] = OpenAPIPath{}
		}
		responses := map[string]OpenAPIResponse{
			strconv.Itoa(op.Response.Status): responseForSchema(op.Response.Ref, statusDescription(op.Response.Status), op.Response.ContentType),
			"default":                        responseForSchema("ErrorEnvelope", "error", ""),
		}
		operation := OpenAPIOperation{
			OperationID:    op.OperationID,
			Summary:        op.Summary,
			Tags:           []string{ir.Context},
			Parameters:     pathParameters(op.Path),
			Responses:      responses,
			Security:       securityForAuth(op.Auth),
			RiidoScopes:    append([]string(nil), op.Auth.Scopes...),
			RiidoRBAC:      op.RBACPolicy,
			RiidoKind:      op.Kind,
			RiidoClient:    deriveClientMeta(op.Client),
			RiidoScenarios: append([]string(nil), op.ScenarioIDs...),
		}
		if op.Request != nil {
			operation.RequestBody = &OpenAPIRequestBody{
				Required: op.Request.Required,
				Content:  jsonContent(op.Request.Ref),
			}
		}
		paths[op.Path][method] = operation
	}
	return OpenAPISpec{
		OpenAPI: OpenAPIVersion,
		Info: OpenAPIInfo{
			Title:   ir.ContractID,
			Version: ir.Service.SchemaVersion,
		},
		Tags:               []OpenAPITag{{Name: ir.Context}},
		RiidoClientModules: copyClientModules(ir.ClientModules),
		Paths:              paths,
		Components: OpenAPIComponents{
			Schemas:         components,
			SecuritySchemes: securitySchemesForIR(ir),
		},
	}, nil
}
