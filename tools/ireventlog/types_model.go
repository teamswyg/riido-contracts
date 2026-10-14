package main

type model struct {
	Manifest               manifest
	EventCount             int
	TransitionCount        int
	NonTransitionCount     int
	NativeConfigCounts     nativeConfigCounts
	TransitionEvents       []string
	TaskFSMTransitionCount int
	TaskFSMTriggerCount    int
	CanonicalEventFields   int
	ReduceResultFields     int
	ReduceResultFieldNames []string
}
