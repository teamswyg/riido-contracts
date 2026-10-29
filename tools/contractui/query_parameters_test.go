package main

import "testing"

func TestRunReadPaginationSurvivesBundle(t *testing.T) {
	data, err := buildBundle()
	if err != nil {
		t.Fatal(err)
	}
	found := 0
	for _, contract := range data.Contracts {
		for _, operation := range contract.Operations {
			if operation.OperationID != "listAIAgentRunReads" && operation.OperationID != "listAIAgentRunReadsV2" {
				continue
			}
			found++
			if len(operation.QueryParameters) != 1 || operation.QueryParameters[0].Name != "page" || operation.QueryParameters[0].Default == nil || *operation.QueryParameters[0].Default != 1 {
				t.Fatalf("pagination lost for %s", operation.OperationID)
			}
		}
	}
	if found != 2 {
		t.Fatalf("run reads routes = %d", found)
	}
}
