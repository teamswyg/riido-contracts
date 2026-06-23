package main

import "testing"

func TestBuildBundleCoversSaaSContracts(t *testing.T) {
	data, err := buildBundle()
	if err != nil {
		t.Fatalf("buildBundle: %v", err)
	}
	if len(data.Contracts) != 2 {
		t.Fatalf("contracts = %d", len(data.Contracts))
	}
	if totalOperations(data) != 61 {
		t.Fatalf("operation count = %d", totalOperations(data))
	}
	if !schemaHasProperty(data, "DeviceDaemonRecord", "app_version") {
		t.Fatal("DeviceDaemonRecord.app_version must be visible in contract UI bundle")
	}
	if !scenarioExists(data, "figma onboarding runtime selection") {
		t.Fatal("Figma scenarios must be visible in contract UI bundle")
	}
}

func TestGeneratedContractUIBundleIsFresh(t *testing.T) {
	if err := verify(); err != nil {
		t.Fatal(err)
	}
}
