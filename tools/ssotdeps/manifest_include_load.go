package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func loadFactDocument(path string) (fact, error) {
	var doc factDocument
	if err := loadStrictJSON(path, &doc); err != nil {
		return fact{}, err
	}
	if doc.SchemaVersion != factSchemaVersion {
		return fact{}, fmt.Errorf("%s schema_version = %q", path, doc.SchemaVersion)
	}
	return doc.Fact, nil
}

func loadRepoDependencyDocument(path string) (repoDependency, error) {
	var doc repoDependencyDocument
	if err := loadStrictJSON(path, &doc); err != nil {
		return repoDependency{}, err
	}
	if doc.SchemaVersion != repoEdgeSchemaVersion {
		return repoDependency{}, fmt.Errorf("%s schema_version = %q", path, doc.SchemaVersion)
	}
	return doc.RepoDependency, nil
}

func loadStrictJSON(path string, out any) error {
	body, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s: %w", path, err)
	}
	dec := json.NewDecoder(strings.NewReader(string(body)))
	dec.DisallowUnknownFields()
	if err := dec.Decode(out); err != nil {
		return fmt.Errorf("decode %s: %w", path, err)
	}
	var extra struct{}
	if err := dec.Decode(&extra); !errors.Is(err, io.EOF) {
		return fmt.Errorf("decode %s: trailing data", path)
	}
	return nil
}
