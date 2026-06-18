package main

func atom(n node) string {
	if !n.isAtom() {
		return ""
	}
	return n.atom
}

func atomList(n node) []string {
	if n.isAtom() {
		value := atom(n)
		if value == "" {
			return nil
		}
		return []string{value}
	}
	out := make([]string, 0, len(n.list))
	for _, item := range n.list {
		value := atom(item)
		if value != "" {
			out = append(out, value)
		}
	}
	return out
}
