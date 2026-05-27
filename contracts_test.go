package contracts

import "testing"

func TestContractSetIdentity(t *testing.T) {
	if ModulePath != "github.com/teamswyg/riido-contracts" {
		t.Fatalf("ModulePath = %q", ModulePath)
	}
	if ContractSetVersion == "" {
		t.Fatal("ContractSetVersion is empty")
	}
}
