package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/teamswyg/riido-contracts/apicontract"
)

const (
	dslGlob = "apicontract/fixtures/*.dsl.riido.json"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "apicontract:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	command := "verify"
	if len(args) > 0 {
		command = args[0]
	}
	switch command {
	case "generate":
		return generate()
	case "verify":
		return verify()
	default:
		return errors.New("usage: go run ./tools/apicontract [verify|generate]")
	}
}

func generate() error {
	generated, err := generatedFixtures()
	if err != nil {
		return err
	}
	for path, body := range generated {
		if err := os.WriteFile(path, body, 0o644); err != nil {
			return fmt.Errorf("write %s: %w", path, err)
		}
	}
	return nil
}

func verify() error {
	generated, err := generatedFixtures()
	if err != nil {
		return err
	}
	for path, want := range generated {
		got, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read %s: %w", path, err)
		}
		if !bytes.Equal(got, want) {
			return fmt.Errorf("%s drifted; run go run ./tools/apicontract generate", path)
		}
	}
	return nil
}

func generatedFixtures() (map[string][]byte, error) {
	dslPaths, err := filepath.Glob(dslGlob)
	if err != nil {
		return nil, err
	}
	if len(dslPaths) == 0 {
		return nil, fmt.Errorf("no DSL fixtures match %s", dslGlob)
	}
	generated := map[string][]byte{}
	for _, dslPath := range dslPaths {
		dsl, err := loadDSL(dslPath)
		if err != nil {
			return nil, err
		}
		ir, err := apicontract.GenerateIR(dsl)
		if err != nil {
			return nil, err
		}
		openAPI, err := apicontract.GenerateOpenAPI(ir)
		if err != nil {
			return nil, err
		}
		irJSON, err := apicontract.MarshalCanonical(ir)
		if err != nil {
			return nil, fmt.Errorf("marshal IR: %w", err)
		}
		openAPIJSON, err := apicontract.MarshalCanonical(openAPI)
		if err != nil {
			return nil, fmt.Errorf("marshal OpenAPI: %w", err)
		}
		stem := filepath.Clean(dslPath[:len(dslPath)-len(".dsl.riido.json")])
		generated[stem+".ir.riido.json"] = irJSON
		generated[stem+".openapi.json"] = openAPIJSON
	}
	return generated, nil
}

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
