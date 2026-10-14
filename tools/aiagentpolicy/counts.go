package main

import "regexp"

var apiPathPattern = regexp.MustCompile(`\b(GET|POST|DELETE|PATCH|PUT) /v[0-9]/[A-Za-z0-9_./{}-]+`)

func countAPIPaths(doc string) int {
	return len(apiPathPattern.FindAllString(doc, -1))
}
