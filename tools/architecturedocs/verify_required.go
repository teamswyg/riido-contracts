package main

import "os"

func verifyRequiredDocs(root string, m manifest) error {
	for _, rel := range m.RequiredDocs {
		path, err := resolve(root, rel)
		if err != nil {
			return err
		}
		if _, err := os.Stat(path); err != nil {
			return err
		}
	}
	return nil
}
