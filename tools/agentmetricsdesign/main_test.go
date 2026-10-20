package main

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

func TestMetricContractAndGeneratedDocument(t *testing.T) {
	body, err := os.ReadFile("../../" + manifestPath)
	if err != nil {
		t.Fatal(err)
	}
	var d contract
	if err := json.Unmarshal(body, &d); err != nil {
		t.Fatal(err)
	}
	if err := validate(d); err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile("../../" + docPath)
	if err != nil || string(got) != render(d) {
		t.Fatal("generated metric document drift")
	}
	if d.Baseline[0].DaemonMS != 780267 || d.Baseline[1].ProviderMS != 92285 || !strings.Contains(d.Provenance, "사용자 제공") {
		t.Fatal("baseline provenance changed")
	}
	for _, mutate := range []func(*contract){func(d *contract) { d.Unknown = 0 }, func(d *contract) { d.Minimum = 1 }, func(d *contract) { d.Target = 100000 }, func(d *contract) { d.Metrics = nil }} {
		candidate := d
		mutate(&candidate)
		if validate(candidate) == nil {
			t.Fatal("invalid metric boundary accepted")
		}
	}
}
