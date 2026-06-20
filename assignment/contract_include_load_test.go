package assignment

import "testing"

func loadContractIncludes(t *testing.T, contract *executableContract) {
	t.Helper()
	loadContractStateIncludes(t, contract)
	loadContractNamedValueIncludes(t, contract)
	loadContractSingleIncludes(t, contract)
	loadContractPayloadIncludes(t, contract)
}

func contractIncludePath(file string) string {
	return file
}
