package main

type nativeConfigCase struct {
	Key string
	Go  string
}

func nativeConfigGroups(enum enumSpec) map[string][]enumValue {
	groups := map[string][]enumValue{}
	for _, value := range enum.Values {
		requirement := value.Attrs["native-config"]
		if requirement == "" || requirement == "required" {
			continue
		}
		groups[requirement] = append(groups[requirement], value)
	}
	return groups
}

func nativeConfigOrder() []nativeConfigCase {
	return []nativeConfigCase{
		{"pre-execute", "NativeConfigOptionalPreExecute"},
		{"phase-dependent", "NativeConfigPhaseDependent"},
		{"forbidden", "NativeConfigForbidden"},
	}
}
