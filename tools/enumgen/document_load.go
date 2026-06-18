package main

import (
	"fmt"
	"os"
)

func loadDocument(path string) (document, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return document{}, fmt.Errorf("read %s: %w", path, err)
	}
	root, err := parseSExpr(string(body))
	if err != nil {
		return document{}, err
	}
	return documentFromNode(root)
}
