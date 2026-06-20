package main

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

func loadJSON[T any](path string) (T, error) {
	var out T
	file, err := os.Open(path)
	if err != nil {
		return out, err
	}
	defer file.Close()
	dec := json.NewDecoder(file)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&out); err != nil {
		return out, err
	}
	var extra struct{}
	if err := dec.Decode(&extra); !errors.Is(err, io.EOF) {
		return out, errors.New("trailing JSON")
	}
	return out, nil
}

func writeFile(path string, body []byte) error {
	return os.WriteFile(path, body, 0o644)
}
