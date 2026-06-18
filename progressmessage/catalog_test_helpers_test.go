package progressmessage

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"
	"testing"
)

func loadTestDSL(t *testing.T) DSLDocument {
	t.Helper()
	body, err := os.ReadFile("catalog.dsl.riido.json")
	if err != nil {
		t.Fatalf("read DSL: %v", err)
	}
	var dsl DSLDocument
	dec := json.NewDecoder(bytes.NewReader(body))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&dsl); err != nil {
		t.Fatalf("decode DSL: %v", err)
	}
	return dsl
}

func assertFixture(t *testing.T, path string, value any) {
	t.Helper()
	want, err := MarshalCanonical(value)
	if err != nil {
		t.Fatalf("MarshalCanonical: %v", err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("%s drifted; run go run ./tools/progressmessage generate", path)
	}
}

func argNames(args []MessageArg) string {
	names := make([]string, 0, len(args))
	for _, arg := range args {
		names = append(names, arg.Name)
	}
	return strings.Join(names, ",")
}
