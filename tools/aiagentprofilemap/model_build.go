package main

func buildModel(root string, m manifest) (model, error) {
	dsl, err := readLooseJSONFile[contractFixture](resolve(root, m.DSLFixture))
	if err != nil {
		return model{}, err
	}
	ir, err := readLooseJSONFile[contractFixture](resolve(root, m.IRFixture))
	if err != nil {
		return model{}, err
	}
	openapi, err := readLooseJSONFile[openAPIDoc](resolve(root, m.OpenAPIFixture))
	if err != nil {
		return model{}, err
	}
	model := model{Manifest: m, DSLIRMatch: true, OpenAPIMatch: true}
	model.Operation = findOperation(dsl.Operations, m.RequiredOperation.OperationID)
	model.ScenarioCount = len(model.Operation.Scenarios)
	model.DSLIRMatch = sameOperationTuple(model.Operation, findOperation(ir.Operations, m.RequiredOperation.OperationID))
	model.OpenAPIMatch = sameOpenAPIOperation(openapi, m.RequiredOperation)
	model.Schemas = selectSchemas(dsl.Schemas, m.RequiredSchemaFields)
	model.Policy = findPolicy(dsl.Policies, m.PolicyID)
	model.MapShapePass = assignedProfileMapShapePass(model.Schemas)
	return model, nil
}
