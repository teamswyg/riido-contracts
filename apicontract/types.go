package apicontract

const (
	DSLSchemaVersion = "riido-api-dsl.v1"
	IRSchemaVersion  = "riido-api-ir.v1"
	OpenAPIVersion   = "3.1.0"
)

type DSLDocument struct {
	SchemaVersion string         `json:"schema_version"`
	ContractID    string         `json:"contract_id"`
	Context       string         `json:"context"`
	Service       Service        `json:"service"`
	ClientModules []ClientModule `json:"client_modules,omitempty"`
	Resources     []Resource     `json:"resources,omitempty"`
	Policies      []Policy       `json:"policies,omitempty"`
	Enums         []Enum         `json:"enums,omitempty"`
	SumTypes      []SumType      `json:"sum_types,omitempty"`
	Schemas       []Schema       `json:"schemas"`
	Operations    []DSLOperation `json:"operations"`
}

type Service struct {
	Name          string `json:"name"`
	SchemaVersion string `json:"schema_version"`
}

type ClientModule struct {
	Module      string            `json:"module"`
	Description string            `json:"description,omitempty"`
	Namespaces  []ClientNamespace `json:"namespaces,omitempty"`
}

type ClientNamespace struct {
	Path        []string `json:"path"`
	Description string   `json:"description,omitempty"`
}

type Resource struct {
	Name       string           `json:"name"`
	OwnerField string           `json:"owner_field,omitempty"`
	Visibility []VisibilityRule `json:"visibility,omitempty"`
}

type VisibilityRule struct {
	Name  string   `json:"name"`
	Read  []string `json:"read,omitempty"`
	Write []string `json:"write,omitempty"`
}

type Policy struct {
	PolicyID    string   `json:"policy_id"`
	Kind        string   `json:"kind"`
	Description string   `json:"description"`
	Rules       []string `json:"rules,omitempty"`
}

type Enum struct {
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	Description string      `json:"description,omitempty"`
	Values      []EnumValue `json:"values"`
}

type EnumValue struct {
	Value       string `json:"value"`
	Description string `json:"description,omitempty"`
}

type SumType struct {
	Name          string           `json:"name"`
	Discriminator string           `json:"discriminator"`
	Description   string           `json:"description,omitempty"`
	Variants      []SumTypeVariant `json:"variants"`
}

type SumTypeVariant struct {
	Kind        string `json:"kind"`
	Schema      string `json:"schema"`
	Description string `json:"description,omitempty"`
}

type Schema struct {
	Name        string     `json:"name"`
	Description string     `json:"description,omitempty"`
	Type        string     `json:"type"`
	Required    []string   `json:"required,omitempty"`
	Properties  []Property `json:"properties,omitempty"`
}

type Property struct {
	Name                 string    `json:"name,omitempty"`
	Type                 string    `json:"type,omitempty"`
	Description          string    `json:"description,omitempty"`
	Format               string    `json:"format,omitempty"`
	MaxLength            *int      `json:"max_length,omitempty"`
	Enum                 []string  `json:"enum,omitempty"`
	Ref                  string    `json:"ref,omitempty"`
	Items                *Property `json:"items,omitempty"`
	AdditionalProperties bool      `json:"additional_properties,omitempty"`
}

type DSLOperation struct {
	OperationID string      `json:"operation_id"`
	Kind        string      `json:"kind"`
	Summary     string      `json:"summary"`
	Description string      `json:"description,omitempty"`
	Method      string      `json:"method"`
	Path        string      `json:"path"`
	Resource    string      `json:"resource"`
	Action      string      `json:"action"`
	Client      *ClientMeta `json:"client,omitempty"`
	Auth        Auth        `json:"auth"`
	RBACPolicy  string      `json:"rbac_policy,omitempty"`
	Request     *MessageRef `json:"request,omitempty"`
	Response    ResponseRef `json:"response"`
	Scenarios   []Scenario  `json:"scenarios,omitempty"`
}

type ClientMeta struct {
	Module        string   `json:"module"`
	FacadePath    []string `json:"facade_path"`
	GeneratedPath string   `json:"generated_path,omitempty"`
	CacheTag      string   `json:"cache_tag,omitempty"`
	Invalidates   []string `json:"invalidates,omitempty"`
}

type Auth struct {
	Scheme string   `json:"scheme"`
	Header string   `json:"header,omitempty"`
	Scopes []string `json:"scopes,omitempty"`
}

type MessageRef struct {
	Ref      string `json:"ref"`
	Required bool   `json:"required,omitempty"`
}

type ResponseRef struct {
	Status      int    `json:"status"`
	Ref         string `json:"ref"`
	ContentType string `json:"content_type,omitempty"`
}

type Scenario struct {
	Name  string `json:"name"`
	Given string `json:"given"`
	When  string `json:"when"`
	Then  string `json:"then"`
}

type IRDocument struct {
	SchemaVersion       string         `json:"schema_version"`
	ContractID          string         `json:"contract_id"`
	SourceSchemaVersion string         `json:"source_schema_version"`
	Context             string         `json:"context"`
	Service             Service        `json:"service"`
	ClientModules       []ClientModule `json:"client_modules,omitempty"`
	Resources           []Resource     `json:"resources,omitempty"`
	Policies            []Policy       `json:"policies,omitempty"`
	Enums               []Enum         `json:"enums,omitempty"`
	SumTypes            []SumType      `json:"sum_types,omitempty"`
	Components          []IRComponent  `json:"components"`
	Operations          []IROperation  `json:"operations"`
	Scenarios           []IRScenario   `json:"scenarios,omitempty"`
}

type IRComponent struct {
	Name   string `json:"name"`
	Schema Schema `json:"schema"`
}

type IROperation struct {
	OperationID string      `json:"operation_id"`
	Kind        string      `json:"kind"`
	Method      string      `json:"method"`
	Path        string      `json:"path"`
	Resource    string      `json:"resource"`
	Action      string      `json:"action"`
	Client      *ClientMeta `json:"client,omitempty"`
	Summary     string      `json:"summary"`
	Auth        Auth        `json:"auth"`
	RBACPolicy  string      `json:"rbac_policy,omitempty"`
	Request     *MessageRef `json:"request,omitempty"`
	Response    ResponseRef `json:"response"`
	ScenarioIDs []string    `json:"scenario_ids,omitempty"`
}

type IRScenario struct {
	ScenarioID  string `json:"scenario_id"`
	OperationID string `json:"operation_id"`
	Name        string `json:"name"`
	Given       string `json:"given"`
	When        string `json:"when"`
	Then        string `json:"then"`
}

type OpenAPISpec struct {
	OpenAPI            string                 `json:"openapi"`
	Info               OpenAPIInfo            `json:"info"`
	Tags               []OpenAPITag           `json:"tags,omitempty"`
	RiidoClientModules []ClientModule         `json:"x-riido-client-modules,omitempty"`
	Paths              map[string]OpenAPIPath `json:"paths"`
	Components         OpenAPIComponents      `json:"components"`
}

type OpenAPIInfo struct {
	Title   string `json:"title"`
	Version string `json:"version"`
}

type OpenAPITag struct {
	Name string `json:"name"`
}

type OpenAPIPath map[string]OpenAPIOperation

type OpenAPIOperation struct {
	OperationID    string                     `json:"operationId"`
	Summary        string                     `json:"summary,omitempty"`
	Tags           []string                   `json:"tags,omitempty"`
	Parameters     []OpenAPIParameter         `json:"parameters,omitempty"`
	RequestBody    *OpenAPIRequestBody        `json:"requestBody,omitempty"`
	Responses      map[string]OpenAPIResponse `json:"responses"`
	Security       []map[string][]string      `json:"security,omitempty"`
	RiidoScopes    []string                   `json:"x-riido-auth-scopes,omitempty"`
	RiidoRBAC      string                     `json:"x-riido-rbac-policy,omitempty"`
	RiidoKind      string                     `json:"x-riido-operation-kind,omitempty"`
	RiidoClient    *ClientMeta                `json:"x-riido-client,omitempty"`
	RiidoScenarios []string                   `json:"x-riido-scenarios,omitempty"`
}

type OpenAPIParameter struct {
	Name     string         `json:"name"`
	In       string         `json:"in"`
	Required bool           `json:"required"`
	Schema   map[string]any `json:"schema"`
}

type OpenAPIRequestBody struct {
	Required bool                    `json:"required,omitempty"`
	Content  map[string]OpenAPIMedia `json:"content"`
}

type OpenAPIResponse struct {
	Description string                  `json:"description"`
	Content     map[string]OpenAPIMedia `json:"content,omitempty"`
}

type OpenAPIMedia struct {
	Schema map[string]any `json:"schema"`
}

type OpenAPIComponents struct {
	Schemas         map[string]map[string]any        `json:"schemas"`
	SecuritySchemes map[string]OpenAPISecurityScheme `json:"securitySchemes,omitempty"`
}

type OpenAPISecurityScheme struct {
	Type         string `json:"type"`
	In           string `json:"in,omitempty"`
	Name         string `json:"name,omitempty"`
	Scheme       string `json:"scheme,omitempty"`
	BearerFormat string `json:"bearerFormat,omitempty"`
}
