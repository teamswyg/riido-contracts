package main

import "fmt"

func verifyModel(model model) error {
	if err := verifyCounts(model); err != nil {
		return err
	}
	if err := verifyTaskTriggers(); err != nil {
		return err
	}
	return verifyReducerSurface(model)
}

func verifyCounts(model model) error {
	m := model.Manifest
	checks := map[string][2]int{
		"event count":          {model.EventCount, m.ExpectedEventCount},
		"transition count":     {model.TransitionCount, m.ExpectedTransitionCount},
		"non-transition count": {model.NonTransitionCount, m.ExpectedNonTransitionCount},
		"task fsm transitions": {model.TaskFSMTransitionCount, m.ExpectedTaskFSMTransitionCount},
		"task fsm triggers":    {model.TaskFSMTriggerCount, m.ExpectedTaskFSMTriggerCount},
		"canonical fields":     {model.CanonicalEventFields, m.ExpectedCanonicalEventFields},
		"reduce result fields": {model.ReduceResultFields, m.ExpectedReduceResultFields},
		"ncv forbidden":        {model.NativeConfigCounts.Forbidden, m.ExpectedNativeConfigCounts.Forbidden},
		"ncv pre-execute":      {model.NativeConfigCounts.PreExecute, m.ExpectedNativeConfigCounts.PreExecute},
		"ncv required":         {model.NativeConfigCounts.Required, m.ExpectedNativeConfigCounts.Required},
		"ncv phase-dependent":  {model.NativeConfigCounts.PhaseDependent, m.ExpectedNativeConfigCounts.PhaseDependent},
	}
	for name, values := range checks {
		if values[0] != values[1] {
			return fmt.Errorf("%s = %d, want %d", name, values[0], values[1])
		}
	}
	return nil
}
