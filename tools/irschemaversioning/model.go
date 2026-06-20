package main

import (
	"reflect"

	"github.com/teamswyg/riido-contracts/ir"
)

func buildModel(m manifest) model {
	return model{
		Manifest:                 m,
		CanonicalEventFields:     reflect.TypeOf(ir.CanonicalEvent{}).NumField(),
		EventScopeCount:          len(eventScopes()),
		CommonRequiredCount:      len(commonRequiredFields()),
		ActorIDConditional:       1,
		RunRequiredFieldCount:    len(runRequiredFields()),
		FakePlaceholderFields:    len(fakePlaceholderFields()),
		FakePlaceholderValues:    len(fakePlaceholderValues()),
		ViolationCodeCount:       len(violationCodes()),
		NativeConfigClassCount:   len(nativeConfigClasses()),
		RunContextFields:         reflect.TypeOf(ir.RunContext{}).NumField(),
		ValidateEntrypoints:      2,
		ScopeRules:               scopeRules(),
		CanonicalEventFieldNames: fieldNames(reflect.TypeOf(ir.CanonicalEvent{})),
		RunContextFieldNames:     fieldNames(reflect.TypeOf(ir.RunContext{})),
	}
}

func fieldNames(t reflect.Type) []string {
	out := make([]string, t.NumField())
	for i := range t.NumField() {
		out[i] = t.Field(i).Name
	}
	return out
}
