package main

import (
	"bytes"
	"fmt"
	"os"
)

func verifyDoc(path, expected string) error {
	current, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if !bytes.Equal(current, []byte(expected)) {
		return fmt.Errorf("%s is stale; run go run ./tools/aiagentpolicy -write-doc", path)
	}
	return nil
}
