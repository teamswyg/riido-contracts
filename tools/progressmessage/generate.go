package main

import "os"

func generate() error {
	body, err := generatedIR()
	if err != nil {
		return err
	}
	return os.WriteFile(irPath, body, 0o644)
}
