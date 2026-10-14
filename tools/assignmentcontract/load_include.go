package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

func loadInclude(path, label, want string, dest any) error {
	body, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	dec := json.NewDecoder(bytes.NewReader(body))
	dec.DisallowUnknownFields()
	if err := dec.Decode(dest); err != nil {
		return fmt.Errorf("decode %s: %w", label, err)
	}
	if err := dec.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return fmt.Errorf("decode %s: trailing data", label)
	}
	if got := includeSchema(dest); got != want {
		return fmt.Errorf("%s schema_version = %q, want %q", label, got, want)
	}
	return nil
}

func includeSchema(dest any) string {
	switch doc := dest.(type) {
	case *stateDocument:
		return doc.SchemaVersion
	case *namedValueDocument:
		return doc.SchemaVersion
	case *executionIDDocument:
		return doc.SchemaVersion
	case *approvalContractDocument:
		return doc.SchemaVersion
	case *payloadFieldDocument:
		return doc.SchemaVersion
	default:
		return ""
	}
}
