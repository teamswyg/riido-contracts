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
		ClientModules:       copyClientModules(dsl.ClientModules),
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
			Client:      deriveClientMeta(op.Client),
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
	clientModules, err := validateClientModules(dsl.ClientModules)
	if err != nil {
		return err
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
			if property.MaxLength != nil && *property.MaxLength <= 0 {
				return fmt.Errorf("apicontract: schema %q property %q max_length must be positive", schema.Name, property.Name)
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
	cacheTags := map[string]string{}
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
		if err := validateAuth(op.OperationID, op.Auth); err != nil {
			return err
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
		if len(clientModules) > 0 {
			if op.Client == nil {
				return fmt.Errorf("apicontract: operation %q missing client metadata", op.OperationID)
			}
			if err := validateClientMeta(op.OperationID, strings.ToUpper(op.Method), *op.Client, clientModules); err != nil {
				return err
			}
			if strings.EqualFold(op.Method, "GET") {
				if prev, exists := cacheTags[op.Client.CacheTag]; exists {
					return fmt.Errorf("apicontract: duplicate client cache_tag %q on %s and %s", op.Client.CacheTag, prev, op.OperationID)
				}
				cacheTags[op.Client.CacheTag] = op.OperationID
			}
		} else if op.Client != nil {
			return fmt.Errorf("apicontract: operation %q declares client metadata without client_modules", op.OperationID)
		}
	}
	if len(clientModules) > 0 {
		for _, op := range dsl.Operations {
			if op.Client == nil {
				continue
			}
			for _, tag := range op.Client.Invalidates {
				if _, ok := cacheTags[tag]; !ok {
					return fmt.Errorf("apicontract: operation %q invalidates unknown client cache_tag %q", op.OperationID, tag)
				}
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
		if err := validateAuth(op.OperationID, op.Auth); err != nil {
			return err
		}
		if _, ok := components[op.Response.Ref]; !ok {
			return fmt.Errorf("apicontract: IR operation %q response schema %q is missing", op.OperationID, op.Response.Ref)
		}
		if op.Request != nil {
			if _, ok := components[op.Request.Ref]; !ok {
				return fmt.Errorf("apicontract: IR operation %q request schema %q is missing", op.OperationID, op.Request.Ref)
			}
		}
	}
	if len(ir.ClientModules) > 0 {
		clientModules, err := validateClientModules(ir.ClientModules)
		if err != nil {
			return err
		}
		cacheTags := map[string]string{}
		for _, op := range ir.Operations {
			if op.Client == nil {
				return fmt.Errorf("apicontract: IR operation %q missing client metadata", op.OperationID)
			}
			if err := validateClientMeta(op.OperationID, op.Method, *op.Client, clientModules); err != nil {
				return err
			}
			if strings.EqualFold(op.Method, "GET") {
				if prev, exists := cacheTags[op.Client.CacheTag]; exists {
					return fmt.Errorf("apicontract: duplicate IR client cache_tag %q on %s and %s", op.Client.CacheTag, prev, op.OperationID)
				}
				cacheTags[op.Client.CacheTag] = op.OperationID
			}
		}
		for _, op := range ir.Operations {
			for _, tag := range op.Client.Invalidates {
				if _, ok := cacheTags[tag]; !ok {
					return fmt.Errorf("apicontract: IR operation %q invalidates unknown client cache_tag %q", op.OperationID, tag)
				}
			}
		}
	}
	return nil
}

func validateAuth(operationID string, auth Auth) error {
	switch auth.Scheme {
	case "bearer":
		if strings.TrimSpace(auth.Header) != "" {
			return fmt.Errorf("apicontract: operation %q bearer auth must not set header", operationID)
		}
	case "apiKey":
		if strings.TrimSpace(auth.Header) == "" {
			return fmt.Errorf("apicontract: operation %q apiKey auth must set header", operationID)
		}
	default:
		return fmt.Errorf("apicontract: operation %q has unsupported auth scheme %q", operationID, auth.Scheme)
	}
	return nil
}

func validateClientModules(modules []ClientModule) (map[string]struct{}, error) {
	out := map[string]struct{}{}
	for _, module := range modules {
		name := strings.TrimSpace(module.Module)
		if name == "" {
			return nil, errors.New("apicontract: client module name is required")
		}
		if _, exists := out[name]; exists {
			return nil, fmt.Errorf("apicontract: duplicate client module %q", name)
		}
		out[name] = struct{}{}
		for _, namespace := range module.Namespaces {
			if len(namespace.Path) == 0 {
				return nil, fmt.Errorf("apicontract: client module %q has empty namespace path", name)
			}
			for _, segment := range namespace.Path {
				if strings.TrimSpace(segment) == "" {
					return nil, fmt.Errorf("apicontract: client module %q has blank namespace path segment", name)
				}
			}
		}
	}
	return out, nil
}

func validateClientMeta(operationID, method string, meta ClientMeta, modules map[string]struct{}) error {
	if _, ok := modules[meta.Module]; !ok {
		return fmt.Errorf("apicontract: operation %q references unknown client module %q", operationID, meta.Module)
	}
	if len(meta.FacadePath) == 0 {
		return fmt.Errorf("apicontract: operation %q missing client facade_path", operationID)
	}
	for _, segment := range meta.FacadePath {
		if strings.TrimSpace(segment) == "" {
			return fmt.Errorf("apicontract: operation %q has blank client facade_path segment", operationID)
		}
	}
	if strings.EqualFold(method, "GET") && strings.TrimSpace(meta.CacheTag) == "" {
		return fmt.Errorf("apicontract: operation %q missing client cache_tag", operationID)
	}
	if generatedPath := strings.TrimSpace(meta.GeneratedPath); generatedPath != "" {
		want := generatedClientPath(meta)
		if generatedPath != want {
			return fmt.Errorf("apicontract: operation %q has client generated_path %q, want %q", operationID, generatedPath, want)
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

func refSchema(name string) map[string]any {
	return map[string]any{"$ref": "#/components/schemas/" + name}
}

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
		out := refSchema(property.Ref)
		if property.Description != "" {
			out["description"] = property.Description
		}
		return out
	}
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
	if property.Items != nil {
		out["items"] = propertyToOpenAPI(*property.Items)
	}
	if property.AdditionalProperties {
		out["additionalProperties"] = true
	}
	return out
}

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

func copyClientModules(modules []ClientModule) []ClientModule {
	if len(modules) == 0 {
		return nil
	}
	out := make([]ClientModule, 0, len(modules))
	for _, module := range modules {
		copied := ClientModule{
			Module:      module.Module,
			Description: module.Description,
			Namespaces:  make([]ClientNamespace, 0, len(module.Namespaces)),
		}
		for _, namespace := range module.Namespaces {
			copied.Namespaces = append(copied.Namespaces, ClientNamespace{
				Path:        append([]string(nil), namespace.Path...),
				Description: namespace.Description,
			})
		}
		out = append(out, copied)
	}
	return out
}

func copyClientMeta(meta *ClientMeta) *ClientMeta {
	if meta == nil {
		return nil
	}
	out := *meta
	out.FacadePath = append([]string(nil), meta.FacadePath...)
	out.Invalidates = append([]string(nil), meta.Invalidates...)
	return &out
}

func deriveClientMeta(meta *ClientMeta) *ClientMeta {
	out := copyClientMeta(meta)
	if out == nil {
		return nil
	}
	out.GeneratedPath = generatedClientPath(*out)
	return out
}

func generatedClientPath(meta ClientMeta) string {
	if strings.TrimSpace(meta.Module) == "" || len(meta.FacadePath) == 0 {
		return ""
	}
	return meta.Module + "." + strings.Join(meta.FacadePath, ".")
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
