package main

import (
	"reflect"
	"strings"

	"github.com/teamswyg/riido-contracts/assignment"
)

func bindingJSONFields() []string {
	t := reflect.TypeOf(assignment.AgentRuntimeBinding{})
	fields := make([]string, 0, t.NumField())
	for i := range t.NumField() {
		tag := t.Field(i).Tag.Get("json")
		name, _, _ := strings.Cut(tag, ",")
		fields = append(fields, name)
	}
	return fields
}
