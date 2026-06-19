package main

import "strings"

func containsPhrase(values []string, phrase string) bool {
	for _, value := range values {
		if strings.Contains(value, phrase) {
			return true
		}
	}
	return false
}

func verifyRenderedDoc(root string, m manifest) error {
	current, err := readLocalRef(root, m.GeneratedDoc)
	if err != nil {
		return err
	}
	expected := renderManifest(m)
	if current != expected {
		return errOutOfDate(m.GeneratedDoc)
	}
	return nil
}
