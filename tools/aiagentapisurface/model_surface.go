package main

func splitSurface(ops []operationTuple) ([]operationTuple, []operationTuple, []operationTuple) {
	var v1, v2 []operationTuple
	for _, op := range ops {
		if isV1(op.Path) {
			v1 = append(v1, op)
		}
		if isV2(op.Path) {
			v2 = append(v2, op)
		}
	}
	return v1, v2, v2Only(v1, v2)
}

func isV1(path string) bool {
	return len(path) >= 4 && path[:4] == "/v1/"
}

func isV2(path string) bool {
	return len(path) >= 4 && path[:4] == "/v2/"
}
