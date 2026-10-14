package main

import "fmt"

func errOutOfDate(path string) error {
	return fmt.Errorf("%s is out of date; regenerate with go run ./tools/architecturedocs write-doc", path)
}
