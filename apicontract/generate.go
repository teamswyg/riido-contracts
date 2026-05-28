package apicontract

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var pathParamPattern = regexp.MustCompile(`\{([^}/]+)\}`)

func GenerateIR(dsl DSLDocument) (IRDocument, error) {
	if err := validateDSL(dsl); err != nil {
		return IRDocument{}, err
	}
	ir := IRDocument{
		SchemaVersion:       IRSchemaVersion,
		ContractID:          dsl.ContractID,
		SourceSchemaVersion: dsl.SchemaVersion,
		Context:             dsl.Context,
		Service:             dsl.Service,
		Resources:           append([]Resource(nil), dsl.Resources...),
		Policies:            append([]Policy(nil), dsl.Policies...),
		Components:          make([]IRComponent, 0, len(dsl.Schemas)),
		Operations:          make([]IROperation, 0, len(dsl.Operations)),
	}
	for _, schema := range dsl.Schemas {
		ir.Components = append(ir.Components, IRComponent{Name: schema.Name, Schema: schema})
	}
	for _, op := range dsl.Operations {
		scenarioIDs := make([]string, 0, len(op.Scenarios))
		for _, scenario := range op.Scenarios {
			id := op.OperationID + "." + slugID(scenario.Name)
			scenarioIDs = append(scenarioIDs, id)
			ir.Scenarios = append(ir.Scenarios, IRScenario{
				ScenarioID:  id,
				OperationID: op.OperationID,
				Name:        scenario.Name,
				Given:       scenario.Given,
				When:        scenario.When,
				Then:        scenario.Then,
			})
		}
		ir.Operations = append(ir.Operations, IROperation{
			OperationID: op.OperationID,
			Kind:        op.Kind,
			Method:      strings.ToUpper(op.Method),
			Path:        op.Path,
			Resource:    op.Resource,
			Action:      op.Action,
			Summary:     op.Summary,
			Auth:        op.Auth,
			RBACPolicy:  op.RBACPolicy,
			Request:     op.Request,
			Response:    op.Response,
			ScenarioIDs: scenarioIDs,
		})
	}
	return ir, nil
}

func GenerateOpenAPI(ir IRDocument) (OpenAPISpec, error) {
	if err := validateIR(ir); err != nil {
		return OpenAPISpec{}, err
	}
	components := map[string]map[string]any{}
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
			strconv.Itoa(op.Response.Status): responseForSchema(op.Response.Ref, statusDescription(op.Response.Status)),
			"default":                        responseForSchema("ErrorEnvelope", "error"),
		}
		operation := OpenAPIOperation{
			OperationID:    op.OperationID,
			Summary:        op.Summary,
			Tags:           []string{ir.Context},
			Parameters:     pathParameters(op.Path),
			Responses:      responses,
			Security:       []map[string][]string{{"bearerAuth": []string{}}},
			RiidoScopes:    append([]string(nil), op.Auth.Scopes...),
			RiidoRBAC:      op.RBACPolicy,
			RiidoKind:      op.Kind,
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
		Tags:  []OpenAPITag{{Name: ir.Context}},
		Paths: paths,
		Components: OpenAPIComponents{
			Schemas: components,
			SecuritySchemes: map[string]OpenAPISecurityScheme{
				"bearerAuth": {Type: "http", Scheme: "bearer", BearerFormat: "opaque"},
			},
		},
	}, nil
}

func MarshalCanonical(value any) ([]byte, error) {
	out, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return nil, err
	}
	return append(out, '\n'), nil
}

func validateDSL(dsl DSLDocument) error {
	if dsl.SchemaVersion != DSLSchemaVersion {
		return fmt.Errorf("apicontract: unsupported DSL schema_version %q", dsl.SchemaVersion)
	}
	for name, value := range map[string]string{
		"contract_id":            dsl.ContractID,
		"context":                dsl.Context,
		"service.name":           dsl.Service.Name,
		"service.schema_version": dsl.Service.SchemaVersion,
	} {
		if strings.TrimSpace(value) == "" {
			return fmt.Errorf("apicontract: %s is required", name)
		}
	}
	schemas := map[string]struct{}{}
	for _, schema := range dsl.Schemas {
		if strings.TrimSpace(schema.Name) == "" {
			return errors.New("apicontract: schema name is required")
		}
		if _, exists := schemas[schema.Name]; exists {
			return fmt.Errorf("apicontract: duplicate schema %q", schema.Name)
		}
		schemas[schema.Name] = struct{}{}
		if schema.Type != "object" {
			return fmt.Errorf("apicontract: schema %q must be object", schema.Name)
		}
		for _, property := range schema.Properties {
			if strings.TrimSpace(property.Name) == "" {
				return fmt.Errorf("apicontract: schema %q has blank property name", schema.Name)
			}
		}
	}
	ops := map[string]struct{}{}
	for _, op := range dsl.Operations {
		if strings.TrimSpace(op.OperationID) == "" {
			return errors.New("apicontract: operation_id is required")
		}
		if _, exists := ops[op.OperationID]; exists {
			return fmt.Errorf("apicontract: duplicate operation %q", op.OperationID)
		}
		ops[op.OperationID] = struct{}{}
		if op.Kind != "query" && op.Kind != "command" {
			return fmt.Errorf("apicontract: operation %q has unsupported kind %q", op.OperationID, op.Kind)
		}
		if !methodAllowed(op.Method) {
			return fmt.Errorf("apicontract: operation %q has unsupported method %q", op.OperationID, op.Method)
		}
		if !strings.HasPrefix(op.Path, "/") {
			return fmt.Errorf("apicontract: operation %q path must start with /", op.OperationID)
		}
		if op.Auth.Scheme != "bearer" {
			return fmt.Errorf("apicontract: operation %q must use bearer auth", op.OperationID)
		}
		if op.Response.Status <= 0 || op.Response.Ref == "" {
			return fmt.Errorf("apicontract: operation %q response is required", op.OperationID)
		}
		if _, ok := schemas[op.Response.Ref]; !ok {
			return fmt.Errorf("apicontract: operation %q response schema %q is missing", op.OperationID, op.Response.Ref)
		}
		if op.Request != nil {
			if _, ok := schemas[op.Request.Ref]; !ok {
				return fmt.Errorf("apicontract: operation %q request schema %q is missing", op.OperationID, op.Request.Ref)
			}
		}
	}
	return nil
}

func validateIR(ir IRDocument) error {
	if ir.SchemaVersion != IRSchemaVersion {
		return fmt.Errorf("apicontract: unsupported IR schema_version %q", ir.SchemaVersion)
	}
	if strings.TrimSpace(ir.ContractID) == "" || strings.TrimSpace(ir.Context) == "" {
		return errors.New("apicontract: IR contract_id and context are required")
	}
	schemas := map[string]struct{}{}
	for _, component := range ir.Components {
		schemas[component.Name] = struct{}{}
	}
	for _, op := range ir.Operations {
		if _, ok := schemas[op.Response.Ref]; !ok {
			return fmt.Errorf("apicontract: IR operation %q response schema %q is missing", op.OperationID, op.Response.Ref)
		}
		if op.Request != nil {
			if _, ok := schemas[op.Request.Ref]; !ok {
				return fmt.Errorf("apicontract: IR operation %q request schema %q is missing", op.OperationID, op.Request.Ref)
			}
		}
	}
	return nil
}

func methodAllowed(method string) bool {
	switch strings.ToUpper(method) {
	case "GET", "POST", "PATCH", "DELETE":
		return true
	default:
		return false
	}
}

func responseForSchema(name, description string) OpenAPIResponse {
	return OpenAPIResponse{
		Description: description,
		Content:     jsonContent(name),
	}
}

func statusDescription(status int) string {
	switch status {
	case 200:
		return "OK"
	case 201:
		return "Created"
	default:
		return "response"
	}
}

func jsonContent(name string) map[string]OpenAPIMedia {
	return map[string]OpenAPIMedia{
		"application/json": {Schema: refSchema(name)},
	}
}

func refSchema(name string) map[string]any {
	return map[string]any{"$ref": "#/components/schemas/" + name}
}

func schemaToOpenAPI(schema Schema) map[string]any {
	out := map[string]any{"type": schema.Type}
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

func propertyToOpenAPI(property Property) map[string]any {
	if property.Ref != "" {
		return refSchema(property.Ref)
	}
	out := map[string]any{}
	if property.Type != "" {
		out["type"] = property.Type
	}
	if property.Format != "" {
		out["format"] = property.Format
	}
	if len(property.Enum) > 0 {
		out["enum"] = append([]string(nil), property.Enum...)
	}
	if property.Items != nil {
		out["items"] = propertyToOpenAPI(*property.Items)
	}
	if property.AdditionalProperties {
		out["additionalProperties"] = true
	}
	return out
}

func pathParameters(path string) []OpenAPIParameter {
	matches := pathParamPattern.FindAllStringSubmatch(path, -1)
	params := make([]OpenAPIParameter, 0, len(matches))
	for _, match := range matches {
		params = append(params, OpenAPIParameter{
			Name:     match[1],
			In:       "path",
			Required: true,
			Schema:   map[string]any{"type": "string"},
		})
	}
	sort.Slice(params, func(i, j int) bool { return params[i].Name < params[j].Name })
	return params
}

func slugID(value string) string {
	value = strings.ToLower(strings.TrimSpace(value))
	var b strings.Builder
	lastDash := false
	for _, r := range value {
		switch {
		case r >= 'a' && r <= 'z':
			b.WriteRune(r)
			lastDash = false
		case r >= '0' && r <= '9':
			b.WriteRune(r)
			lastDash = false
		default:
			if !lastDash && b.Len() > 0 {
				b.WriteByte('-')
				lastDash = true
			}
		}
	}
	return strings.Trim(b.String(), "-")
}
