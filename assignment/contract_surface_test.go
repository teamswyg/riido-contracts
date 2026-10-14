package assignment

import "testing"

func TestAssignmentContractMatchesPackageSurface(t *testing.T) {
	contract := loadContract(t)
	assertContractVersions(t, contract)
	assertAssignmentStatesMatchContract(t, contract)
	assertPollActionsMatchContract(t, contract)
	assertTaskEventsMatchContract(t, contract)
	assertExecutionIdentityContract(t, contract)
	assertApprovalContract(t, contract.ApprovalContract)
	assertAssignmentPayloadFields(t, contract)
}

func assertContractVersions(t *testing.T, contract executableContract) {
	t.Helper()
	if contract.SchemaVersion != ContractSchemaVersion {
		t.Fatalf("contract schema_version = %q, want %q", contract.SchemaVersion, ContractSchemaVersion)
	}
	if contract.ServiceSchemaVersion != SchemaVersion {
		t.Fatalf("service_schema_version = %q, want %q", contract.ServiceSchemaVersion, SchemaVersion)
	}
}
