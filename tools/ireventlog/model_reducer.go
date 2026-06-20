package main

import (
	"reflect"

	"github.com/teamswyg/riido-contracts/ir"
)

func reduceResultFieldNames() []string {
	t := reflect.TypeOf(ir.ReduceResult{})
	out := make([]string, t.NumField())
	for i := range t.NumField() {
		out[i] = t.Field(i).Name
	}
	return out
}
