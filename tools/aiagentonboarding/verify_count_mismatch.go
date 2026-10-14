package main

import "fmt"

func countMismatch(name string, got, want int) error {
	if got != want {
		return fmt.Errorf("%s count = %d, want %d", name, got, want)
	}
	return nil
}
