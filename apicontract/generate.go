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
		Enums:               append([]Enum(nil), dsl.Enums...),
		SumTypes:            append([]SumType(nil), dsl.SumTypes...),
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
	components := map[string]struct{}{}
	for _, enum := range dsl.Enums {
		if strings.TrimSpace(enum.Name) == "" {
			return errors.New("apicontract: enum name is required")
		}
		if enum.Type != "string" {
			return fmt.Errorf("apicontract: enum %q must use string type", enum.Name)
		}
		if len(enum.Values) == 0 {
			return fmt.Errorf("apicontract: enum %q must define values", enum.Name)
		}
		if _, exists := components[enum.Name]; exists {
			return fmt.Errorf("apicontract: duplicate component %q", enum.Name)
		}
		components[enum.Name] = struct{}{}
		values := map[string]struct{}{}
		for _, value := range enum.Values {
			if strings.TrimSpace(value.Value) == "" {
				return fmt.Errorf("apicontract: enum %q has blank value", enum.Name)
			}
			if _, exists := values[value.Value]; exists {
				return fmt.Errorf("apicontract: enum %q has duplicate value %q", enum.Name, value.Value)
			}
			values[value.Value] = struct{}{}
		}
	}
	for _, sumType := range dsl.SumTypes {
		if strings.TrimSpace(sumType.Name) == "" {
			return errors.New("apicontract: sum_type name is required")
		}
		if strings.TrimSpace(sumType.Discriminator) == "" {
			return fmt.Errorf("apicontract: sum_type %q discriminator is required", sumType.Name)
		}
		if len(sumType.Variants) == 0 {
			return fmt.Errorf("apicontract: sum_type %q must define variants", sumType.Name)
		}
		if _, exists := components[sumType.Name]; exists {
			return fmt.Errorf("apicontract: duplicate component %q", sumType.Name)
		}
		components[sumType.Name] = struct{}{}
		kinds := map[string]struct{}{}
		for _, variant := range sumType.Variants {
			if strings.TrimSpace(variant.Kind) == "" || strings.TrimSpace(variant.Schema) == "" {
				return fmt.Errorf("apicontract: sum_type %q variant kind and schema are required", sumType.Name)
			}
			if _, exists := kinds[variant.Kind]; exists {
				return fmt.Errorf("apicontract: sum_type %q has duplicate variant %q", sumType.Name, variant.Kind)
			}
			kinds[variant.Kind] = struct{}{}
		}
	}
	schemas := map[string]struct{}{}
	for _, schema := range dsl.Schemas {
		if strings.TrimSpace(schema.Name) == "" {
			return errors.New("apicontract: schema name is required")
		}
		if _, exists := components[schema.Name]; exists {
			return fmt.Errorf("apicontract: duplicate component %q", schema.Name)
		}
		schemas[schema.Name] = struct{}{}
		components[schema.Name] = struct{}{}
		if schema.Type != "object" {
			return fmt.Errorf("apicontract: schema %q must be object", schema.Name)
		}
		for _, property := range schema.Properties {
			if strings.TrimSpace(property.Name) == "" {
				return fmt.Errorf("apicontract: schema %q has blank property name", schema.Name)
			}
			if err := validatePropertyRef(schema.Name, property, components); err != nil {
				return err
			}
		}
	}
	for _, sumType := range dsl.SumTypes {
		for _, variant := range sumType.Variants {
			if _, ok := schemas[variant.Schema]; !ok {
				return fmt.Errorf("apicontract: sum_type %q variant schema %q is missing", sumType.Name, variant.Schema)
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
		if _, ok := components[op.Response.Ref]; !ok {
			return fmt.Errorf("apicontract: operation %q response schema %q is missing", op.OperationID, op.Response.Ref)
		}
		if op.Request != nil {
			if _, ok := components[op.Request.Ref]; !ok {
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
	components := map[string]struct{}{}
	schemas := map[string]struct{}{}
	for _, enum := range ir.Enums {
		components[enum.Name] = struct{}{}
	}
	for _, sumType := range ir.SumTypes {
		components[sumType.Name] = struct{}{}
	}
	for _, component := range ir.Components {
		schemas[component.Name] = struct{}{}
		components[component.Name] = struct{}{}
	}
	for _, sumType := range ir.SumTypes {
		for _, variant := range sumType.Variants {
			if _, ok := schemas[variant.Schema]; !ok {
				return fmt.Errorf("apicontract: IR sum_type %q variant schema %q is missing", sumType.Name, variant.Schema)
			}
		}
	}
	for _, op := range ir.Operations {
		if _, ok := components[op.Response.Ref]; !ok {
			return fmt.Errorf("apicontract: IR operation %q response schema %q is missing", op.OperationID, op.Response.Ref)
		}
		if op.Request != nil {
			if _, ok := components[op.Request.Ref]; !ok {
				return fmt.Errorf("apicontract: IR operation %q request schema %q is missing", op.OperationID, op.Request.Ref)
			}
		}
	}
	return nil
}

func validatePropertyRef(schemaName string, property Property, components map[string]struct{}) error {
	if property.Ref != "" {
		if _, ok := components[property.Ref]; !ok {
			return fmt.Errorf("apicontract: schema %q property %q ref %q is missing", schemaName, property.Name, property.Ref)
		}
	}
	if property.Items != nil {
		return validatePropertyRef(schemaName, *property.Items, components)
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
