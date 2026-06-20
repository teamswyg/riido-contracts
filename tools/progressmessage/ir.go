package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/teamswyg/riido-contracts/progressmessage"
)

func generatedIR(root string) ([]byte, error) {
	_, _, body, err := buildIR(root)
	return body, err
}

func buildIR(root string) (progressmessage.DSLDocument, progressmessage.IRDocument, []byte, error) {
	data, err := os.ReadFile(resolve(root, dslPath))
	if err != nil {
		return progressmessage.DSLDocument{}, progressmessage.IRDocument{}, nil, fmt.Errorf("read %s: %w", dslPath, err)
	}
	dsl, err := decodeDSL(data)
	if err != nil {
		return progressmessage.DSLDocument{}, progressmessage.IRDocument{}, nil, err
	}
	ir, err := progressmessage.GenerateIR(dsl)
	if err != nil {
		return progressmessage.DSLDocument{}, progressmessage.IRDocument{}, nil, err
	}
	body, err := progressmessage.MarshalCanonical(ir)
	return dsl, ir, body, err
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
