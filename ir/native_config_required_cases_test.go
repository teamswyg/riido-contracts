package ir

func nativeConfigRequiredCases() []nativeConfigRequirementCase {
	return []nativeConfigRequirementCase{
		{EventNativeConfigInjected, NativeConfigRequired},
		{EventRunStarted, NativeConfigRequired},
		{EventTextDelta, NativeConfigRequired},
		{EventValidationPassed, NativeConfigRequired},
		{EventHumanApproved, NativeConfigRequired},
	}
}
