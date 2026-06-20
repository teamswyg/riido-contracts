package main

import "sort"

func operationTuples(ops []operation) []operationTuple {
	out := make([]operationTuple, 0, len(ops))
	for _, op := range ops {
		out = append(out, operationTuple{Method: op.Method, Path: op.Path, OperationID: op.OperationID})
	}
	sortTuples(out)
	return out
}

func sortTuples(values []operationTuple) {
	sort.Slice(values, func(i, j int) bool {
		if values[i].Path != values[j].Path {
			return values[i].Path < values[j].Path
		}
		if values[i].Method != values[j].Method {
			return values[i].Method < values[j].Method
		}
		return values[i].OperationID < values[j].OperationID
	})
}
