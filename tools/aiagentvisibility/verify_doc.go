package main

import (
	"bytes"
	"fmt"
	"os"
)

func verifyDoc(path, expected string) error {
	actual, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if !bytes.Equal(actual, []byte(expected)) {
		return fmt.Errorf("%s is stale; run go run ./tools/aiagentvisibility -write-doc", path)
	}
	return nil
}
