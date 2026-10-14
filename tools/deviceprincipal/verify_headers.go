package main

import "fmt"

func verifyHeaderBoundary(model model) error {
	for _, daemonHeader := range model.DaemonHeaders {
		if contains(model.ClientHeaders, daemonHeader) {
			return fmt.Errorf("daemon header %q overlaps client credential header", daemonHeader)
		}
	}
	return nil
}
