package main

import (
	"fmt"
	"regexp"
)

var idPattern = regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)

func requireID(label, id string) error {
	if !idPattern.MatchString(id) {
		return fmt.Errorf("%s must match %s", label, idPattern.String())
	}
	return nil
}
