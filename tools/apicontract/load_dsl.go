package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/teamswyg/riido-contracts/apicontract"
)

func loadDSL(path string) (apicontract.DSLDocument, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return apicontract.DSLDocument{}, fmt.Errorf("read %s: %w", path, err)
	}
	var dsl apicontract.DSLDocument
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&dsl); err != nil {
		return apicontract.DSLDocument{}, fmt.Errorf("decode %s: %w", path, err)
	}
	var extra struct{}
	if err := dec.Decode(&extra); err != io.EOF {
		return apicontract.DSLDocument{}, fmt.Errorf("decode %s: trailing data", path)
	}
	return dsl, nil
}
