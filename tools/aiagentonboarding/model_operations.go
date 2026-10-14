package main

func addOperations(
	model model,
	dsl contractFixture,
	ir contractFixture,
	openapi openAPIDoc,
) model {
	model.DSLIRMatch = true
	model.OpenAPIMatch = true
	for _, exp := range model.Manifest.RequiredOperations {
		dslOp := findOperation(dsl.Operations, exp.OperationID)
		irOp := findOperation(ir.Operations, exp.OperationID)
		model.Operations = append(model.Operations, dslOp)
		if isOnboardingOperation(dslOp) {
			model.OnboardingOperations = append(model.OnboardingOperations, dslOp)
		} else {
			model.DirectCreateOperations = append(model.DirectCreateOperations, dslOp)
		}
		model.ScenarioCount += len(dslOp.Scenarios)
		model.DSLIRMatch = model.DSLIRMatch && sameOperationTuple(dslOp, irOp)
		model.OpenAPIMatch = model.OpenAPIMatch && openAPIMatches(openapi, dslOp)
	}
	return model
}

func findOperation(ops []operation, id string) operation {
	for _, op := range ops {
		if op.OperationID == id {
			return op
		}
	}
	return operation{}
}
