package main

import (
	"fmt"
	"reflect"
)

func verifyScopeRules(model model) error {
	if !reflect.DeepEqual(model.ScopeRules, model.Manifest.ScopeRules) {
		return fmt.Errorf("scope rules = %#v, want %#v", model.ScopeRules, model.Manifest.ScopeRules)
	}
	return nil
}
