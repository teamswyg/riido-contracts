package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/teamswyg/riido-contracts/progressmessage"
)

func generatedIR() ([]byte, error) {
	data, err := os.ReadFile(dslPath)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", dslPath, err)
	}
	dsl, err := decodeDSL(data)
	if err != nil {
		return nil, err
	}
	ir, err := progressmessage.GenerateIR(dsl)
	if err != nil {
		return nil, err
	}
	return progressmessage.MarshalCanonical(ir)
}

func decodeDSL(data []byte) (progressmessage.DSLDocument, error) {
	var dsl progressmessage.DSLDocument
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&dsl); err != nil {
		return progressmessage.DSLDocument{}, fmt.Errorf("decode %s: %w", dslPath, err)
	}
	var extra struct{}
	if err := dec.Decode(&extra); err != io.EOF {
		return progressmessage.DSLDocument{}, fmt.Errorf("decode %s: trailing data", dslPath)
	}
	return dsl, nil
}
