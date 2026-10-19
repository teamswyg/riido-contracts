package apicontract

import "testing"

func TestGenerateIRRejectsInvalidStringConstraints(t *testing.T) {
	for _, tc := range []struct {
		name     string
		property Property
	}{
		{"negative minimum", Property{Type: "string", MinLength: intPointer(-1)}},
		{"minimum exceeds maximum", Property{Type: "string", MinLength: intPointer(2), MaxLength: intPointer(1)}},
		{"invalid regular expression", Property{Type: "string", Pattern: "["}},
		{"numeric pattern", Property{Type: "integer", Pattern: "^[0-9]+$"}},
		{"reference minimum", Property{Ref: "AgentVisibility", MinLength: intPointer(1)}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			dsl := loadTestDSL(t, "fixtures/control-plane-ai-agent-client.dsl.riido.json")
			tc.property.Name = "invalid_constraint"
			dsl.Schemas[0].Properties = append(dsl.Schemas[0].Properties, tc.property)
			if _, err := GenerateIR(dsl); err == nil {
				t.Fatal("expected invalid string constraint rejection")
			}
		})
	}
}

func TestOpenAPIProjectsZeroMinimumLength(t *testing.T) {
	property := Property{Type: "string", MinLength: intPointer(0), MaxLength: intPointer(128), Pattern: "^[a-z]*$"}
	if err := validateStringConstraints("test", property); err != nil {
		t.Fatal(err)
	}
	projected := propertyToOpenAPI(property)
	if projected["minLength"] != 0 || projected["maxLength"] != 128 || projected["pattern"] != property.Pattern {
		t.Fatalf("string constraint projection = %#v", projected)
	}
}

func intPointer(value int) *int { return &value }
