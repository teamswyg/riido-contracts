package main

import "fmt"

func verifyModel(model model) error {
	if err := verifyCounts(model); err != nil {
		return err
	}
	if err := verifyScopeRules(model); err != nil {
		return err
	}
	if err := verifyValidatorProbes(); err != nil {
		return err
	}
	return nil
}

func verifyCounts(model model) error {
	m := model.Manifest
	checks := map[string][2]int{
		"canonical fields":     {model.CanonicalEventFields, m.ExpectedCanonicalEventFields},
		"scope count":          {model.EventScopeCount, m.ExpectedEventScopeCount},
		"common required":      {model.CommonRequiredCount, m.ExpectedCommonRequiredCount},
		"actor conditional":    {model.ActorIDConditional, m.ExpectedActorIDConditional},
		"run required":         {model.RunRequiredFieldCount, m.ExpectedRunRequiredFieldCount},
		"placeholder fields":   {model.FakePlaceholderFields, m.ExpectedFakePlaceholderFields},
		"placeholder values":   {model.FakePlaceholderValues, m.ExpectedFakePlaceholderValues},
		"violation codes":      {model.ViolationCodeCount, m.ExpectedViolationCodeCount},
		"native config class":  {model.NativeConfigClassCount, m.ExpectedNativeConfigClassCount},
		"run context fields":   {model.RunContextFields, m.ExpectedRunContextFields},
		"validate entrypoints": {model.ValidateEntrypoints, m.ExpectedValidateEntrypoints},
	}
	for name, values := range checks {
		if values[0] != values[1] {
			return fmt.Errorf("%s = %d, want %d", name, values[0], values[1])
		}
	}
	return nil
}
