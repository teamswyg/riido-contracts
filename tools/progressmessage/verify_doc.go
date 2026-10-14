package main

import (
	"bytes"
	"fmt"
	"os"
)

func verifyDoc(path, want string) error {
	got, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if !bytes.Equal(got, []byte(want)) {
		return fmt.Errorf("%s is stale; run go run ./tools/progressmessage verify -write-doc", path)
	}
	return nil
}
