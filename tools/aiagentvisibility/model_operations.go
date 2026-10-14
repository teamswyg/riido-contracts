package main

func addOperations(model model, dsl, ir contractFixture, openapi openAPIDoc) model {
	for _, exp := range model.Manifest.RequiredOperations {
		op := findOperation(dsl.Operations, exp.OperationID)
		if op.OperationID != "" {
			model.Operations = append(model.Operations, op)
			model.ScenarioCount += len(op.Scenarios)
		}
		if !sameOperationTuple(op, findOperation(ir.Operations, exp.OperationID)) {
			model.DSLIRMatch = false
		}
		if !sameOpenAPIOperation(openapi, exp) {
			model.OpenAPIMatch = false
		}
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

func sameOperationTuple(a, b operation) bool {
	if a.OperationID == "" || b.OperationID == "" {
		return false
	}
	return a.Kind == b.Kind && a.Method == b.Method && a.Path == b.Path &&
		responseName(a) == responseName(b) && responseStatus(a) == responseStatus(b) &&
		a.RBACPolicy == b.RBACPolicy
}

func responseName(op operation) string {
	if op.Response == nil {
		return ""
	}
	return op.Response.Ref
}

func responseStatus(op operation) int {
	if op.Response == nil {
		return 0
	}
	return op.Response.Status
}
