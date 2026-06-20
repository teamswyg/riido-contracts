package main

import (
	"reflect"

	"github.com/teamswyg/riido-contracts/ir"
	"github.com/teamswyg/riido-contracts/task"
)

func buildModel(m manifest) model {
	events := ir.AllEventTypes()
	transitions := transitionEvents(events)
	return model{
		Manifest:               m,
		EventCount:             len(events),
		TransitionCount:        len(transitions),
		NonTransitionCount:     len(events) - len(transitions),
		NativeConfigCounts:     countNativeConfig(events),
		TransitionEvents:       eventNames(transitions),
		TaskFSMTransitionCount: len(task.LegalTransitionCodes()),
		TaskFSMTriggerCount:    taskFSMTriggerCount(),
		CanonicalEventFields:   reflect.TypeOf(ir.CanonicalEvent{}).NumField(),
		ReduceResultFields:     reflect.TypeOf(ir.ReduceResult{}).NumField(),
		ReduceResultFieldNames: reduceResultFieldNames(),
	}
}
