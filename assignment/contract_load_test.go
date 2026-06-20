package assignment

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"os"
	"testing"
)

func loadContract(t *testing.T) executableContract {
	t.Helper()
	data, err := os.ReadFile("assignment_contract.riido.json")
	if err != nil {
		t.Fatalf("read assignment contract: %v", err)
	}
	var contract executableContract
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&contract); err != nil {
		t.Fatalf("unmarshal assignment contract: %v", err)
	}
	var trailing struct{}
	if err := decoder.Decode(&trailing); !errors.Is(err, io.EOF) {
		t.Fatal("assignment contract must contain exactly one JSON document")
	}
	loadContractIncludes(t, &contract)
	return contract
}
