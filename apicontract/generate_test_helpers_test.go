package apicontract

import (
	"bytes"
	"encoding/json"
	"os"
	"slices"
	"testing"
)

func loadTestDSL(t *testing.T, path string) DSLDocument {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read DSL fixture: %v", err)
	}
	var dsl DSLDocument
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&dsl); err != nil {
		t.Fatalf("decode DSL fixture: %v", err)
	}
	return dsl
}

func assertFixture(t *testing.T, path string, value any) {
	t.Helper()
	want, err := MarshalCanonical(value)
	if err != nil {
		t.Fatalf("marshal fixture: %v", err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("%s drifted; run go run ./tools/apicontract generate", path)
	}
}

func contains(values []string, needle string) bool {
	return slices.Contains(values, needle)
}
