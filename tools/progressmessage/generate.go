package main

import "os"

func generate() error {
	root, err := resolveRoot(".")
	if err != nil {
		return err
	}
	_, ir, err := buildIR(root)
	if err != nil {
		return err
	}
	files, err := generatedIRFilesFor(ir)
	if err != nil {
		return err
	}
	if err := os.RemoveAll(resolve(root, irMessageDir)); err != nil {
		return err
	}
	if err := writeFile(resolve(root, irPath), files.Root); err != nil {
		return err
	}
	for path, body := range files.Messages {
		if err := writeFile(resolve(root, path), body); err != nil {
			return err
		}
	}
	return nil
}
