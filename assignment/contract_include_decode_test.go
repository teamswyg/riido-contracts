package assignment

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"os"
	"testing"
)

func loadContractInclude(t *testing.T, file, want string, dest any) {
	t.Helper()
	data, err := os.ReadFile(contractIncludePath(file))
	if err != nil {
		t.Fatalf("read contract include %s: %v", file, err)
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(dest); err != nil {
		t.Fatalf("decode contract include %s: %v", file, err)
	}
	var trailing struct{}
	if err := decoder.Decode(&trailing); !errors.Is(err, io.EOF) {
		t.Fatalf("contract include %s must contain exactly one JSON document", file)
	}
	if got := contractIncludeSchema(dest); got != want {
		t.Fatalf("contract include %s schema_version = %q, want %q", file, got, want)
	}
}

func contractIncludeSchema(dest any) string {
	switch doc := dest.(type) {
	case *contractStateDocument:
		return doc.SchemaVersion
	case *contractNamedValueDocument:
		return doc.SchemaVersion
	case *contractExecutionIdentityDocument:
		return doc.SchemaVersion
	case *contractApprovalDocument:
		return doc.SchemaVersion
	case *contractPayloadFieldDocument:
		return doc.SchemaVersion
	default:
		return ""
	}
}
