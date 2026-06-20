package main

import (
	"bytes"
	"os"
)

func verifyDoc(path, expected string) error {
	current, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if !bytes.Equal(current, []byte(expected)) {
		return os.ErrInvalid
	}
	return nil
}
