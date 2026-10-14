package main

func (e enumSpec) hasConst(name string) bool {
	for _, value := range e.Values {
		if value.Const == name {
			return true
		}
	}
	return false
}
