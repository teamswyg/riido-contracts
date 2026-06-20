package main

import "path/filepath"

func resolve(root, slashPath string) string {
	if filepath.IsAbs(slashPath) {
		return slashPath
	}
	return filepath.Join(root, filepath.FromSlash(slashPath))
}
