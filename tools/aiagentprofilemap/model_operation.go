package main

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
		a.RBACPolicy == b.RBACPolicy && clientRoute(a) == clientRoute(b)
}

func clientRoute(op operation) string {
	if op.Client.GeneratedPath != "" {
		return op.Client.GeneratedPath
	}
	return op.Client.CacheTag
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
