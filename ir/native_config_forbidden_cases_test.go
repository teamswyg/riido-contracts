package ir

func nativeConfigForbiddenCases() []nativeConfigRequirementCase {
	return []nativeConfigRequirementCase{
		{EventTaskCreated, NativeConfigForbidden},
		{EventTaskQueued, NativeConfigForbidden},
		{EventRuntimeRegistered, NativeConfigForbidden},
		{EventPolicyBundleLoaded, NativeConfigForbidden},
		{EventUpgradeDetected, NativeConfigForbidden},
	}
}
