package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/teamswyg/riido-contracts/progressmessage"
)

const (
	dslPath = "progressmessage/catalog.dsl.riido.json"
	irPath  = "progressmessage/catalog.ir.riido.json"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "progressmessage:", err)
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
		return errors.New("usage: go run ./tools/progressmessage [verify|generate]")
	}
}

func generate() error {
	body, err := generatedIR()
	if err != nil {
		return err
	}
	return os.WriteFile(irPath, body, 0o644)
}

func verify() error {
	want, err := generatedIR()
	if err != nil {
		return err
	}
	got, err := os.ReadFile(irPath)
	if err != nil {
		return fmt.Errorf("read %s: %w", irPath, err)
	}
	if !bytes.Equal(got, want) {
		return fmt.Errorf("%s drifted; run go run ./tools/progressmessage generate", irPath)
	}
	return nil
}

func generatedIR() ([]byte, error) {
	data, err := os.ReadFile(dslPath)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", dslPath, err)
	}
	var dsl progressmessage.DSLDocument
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&dsl); err != nil {
		return nil, fmt.Errorf("decode %s: %w", dslPath, err)
	}
	var extra struct{}
	if err := dec.Decode(&extra); err != io.EOF {
		return nil, fmt.Errorf("decode %s: trailing data", dslPath)
	}
	ir, err := progressmessage.GenerateIR(dsl)
	if err != nil {
		return nil, err
	}
	return progressmessage.MarshalCanonical(ir)
}
