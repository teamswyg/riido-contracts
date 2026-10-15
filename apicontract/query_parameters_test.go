package apicontract

import "testing"

func TestQueryParametersSurviveIRAndOpenAPI(t *testing.T) {
	dsl := loadTestDSL(t, "fixtures/control-plane-agent-catalog.dsl.riido.json")
	one := 1
	dsl.Operations[0].QueryParameters = []QueryParameter{{Name: "page", Type: "integer", Minimum: &one, Default: &one}}
	ir, err := GenerateIR(dsl)
	if err != nil {
		t.Fatal(err)
	}
	if len(ir.Operations[0].QueryParameters) != 1 {
		t.Fatal("IR lost page")
	}
	spec, err := GenerateOpenAPI(ir)
	if err != nil {
		t.Fatal(err)
	}
	found := false
	for _, methods := range spec.Paths {
		for _, op := range methods {
			for _, p := range op.Parameters {
				if p.Name == "page" {
					found = p.In == "query" && !p.Required && p.Schema["type"] == "integer" && p.Schema["minimum"] == 1 && p.Schema["default"] == 1
				}
			}
		}
	}
	if !found {
		t.Fatal("OpenAPI lost typed/defaulted optional page")
	}
	for _, bad := range [][]QueryParameter{{{Name: "page", Type: "object"}}, {{Name: "page", Type: "integer"}, {Name: "page", Type: "integer"}}, {{Name: "page", Type: "string", Minimum: &one}}} {
		dsl.Operations[0].QueryParameters = bad
		if _, err := GenerateIR(dsl); err == nil {
			t.Fatal("invalid DSL accepted")
		}
		ir.Operations[0].QueryParameters = bad
		if _, err := GenerateOpenAPI(ir); err == nil {
			t.Fatal("invalid IR accepted")
		}
	}
}
