package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"os"
)

func readJSONFile[T any](path string) (T, error) {
	var out T
	body, err := os.ReadFile(path)
	if err != nil {
		return out, err
	}
	dec := json.NewDecoder(bytes.NewReader(body))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&out); err != nil {
		return out, err
	}
	if err := dec.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return out, errors.New("json has trailing data")
	}
	return out, nil
}

func readLooseJSONFile[T any](path string) (T, error) {
	var out T
	body, err := os.ReadFile(path)
	if err != nil {
		return out, err
	}
	if err := json.Unmarshal(body, &out); err != nil {
		return out, err
	}
	return out, nil
}

func writeFile(path string, body []byte) error {
	return os.WriteFile(path, body, 0o644)
}
