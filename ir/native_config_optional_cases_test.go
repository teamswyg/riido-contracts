package ir

func nativeConfigOptionalCases() []nativeConfigRequirementCase {
	return []nativeConfigRequirementCase{
		{EventTaskClaimed, NativeConfigOptionalPreExecute},
		{EventWorkdirPreparing, NativeConfigOptionalPreExecute},
		{EventRuntimePinned, NativeConfigOptionalPreExecute},
		{EventRuntimeHandshakeOK, NativeConfigOptionalPreExecute},
		{EventBlockerRaised, NativeConfigPhaseDependent},
		{EventTaskFailed, NativeConfigPhaseDependent},
		{EventTaskCancelled, NativeConfigPhaseDependent},
		{EventTaskTimedOut, NativeConfigPhaseDependent},
		{EventRuntimePinViolated, NativeConfigPhaseDependent},
		{EventReworkAccepted, NativeConfigPhaseDependent},
	}
}
