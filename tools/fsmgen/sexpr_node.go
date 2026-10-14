package main

type node struct {
	atom string
	list []node
}

func (n node) isAtom() bool {
	return n.list == nil
}
