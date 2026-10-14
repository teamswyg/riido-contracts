package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func readJSON(path string, out any) error {
	body, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s: %w", path, err)
	}
	if err := json.Unmarshal(body, out); err != nil {
		return fmt.Errorf("decode %s: %w", path, err)
	}
	return nil
}

func renderBundleJS(data bundle) ([]byte, error) {
	body, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return nil, err
	}
	var out bytes.Buffer
	out.WriteString(generatedNotice)
	out.WriteString("export const saasContractBundle = ")
	out.Write(body)
	out.WriteString(";\n\nexport default saasContractBundle;\n")
	return out.Bytes(), nil
}
