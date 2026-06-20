package main

const (
	v1Prefix = "/v1/client/"
	v2Prefix = "/v2/client/workspaces/{workspace_id}/"
)

func v2CoversV1(v1, v2 []operationTuple) bool {
	v2Set := normalizedSet(v2)
	for _, op := range v1 {
		if !v2Set[normalizedKey(op)] {
			return false
		}
	}
	return true
}

func v2Only(v1, v2 []operationTuple) []operationTuple {
	v1Set := normalizedSet(v1)
	var out []operationTuple
	for _, op := range v2 {
		if !v1Set[normalizedKey(op)] {
			out = append(out, op)
		}
	}
	return out
}
