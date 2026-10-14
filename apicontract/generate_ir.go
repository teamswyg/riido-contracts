package apicontract

import (
	"strings"
)

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
