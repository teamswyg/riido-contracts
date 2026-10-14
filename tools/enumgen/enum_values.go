package main

func (e enumSpec) valuesWithAttr(attr, want string) []enumValue {
	var out []enumValue
	for _, value := range e.Values {
		if value.Attrs[attr] == want {
			out = append(out, value)
		}
	}
	return out
}

func enumCodeRefs(enum enumSpec, values []enumValue) []string {
	out := make([]string, 0, len(values))
	for _, value := range values {
		out = append(out, enum.codeConst(value.Const))
	}
	return out
}
