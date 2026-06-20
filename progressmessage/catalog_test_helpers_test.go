package progressmessage

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func loadTestDSL(t *testing.T) DSLDocument {
	t.Helper()
	dsl, err := LoadDSL(os.DirFS("."), "catalog.dsl.riido.json")
	if err != nil {
		t.Fatalf("LoadDSL: %v", err)
	}
	return dsl
}

func assertCatalogEqual(t *testing.T, want, got IRDocument) {
	t.Helper()
	want.MessageFiles = nil
	got.MessageFiles = nil
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("catalog does not match generated IR")
	}
}

func argNames(args []MessageArg) string {
	names := make([]string, 0, len(args))
	for _, arg := range args {
		names = append(names, arg.Name)
	}
	return strings.Join(names, ",")
}
