package ir

import "strings"

// fakePlaceholders are sentinel strings that MUST NOT be used to fill an
// identifier field that is conceptually absent.
var fakePlaceholders = map[string]struct{}{
	"unknown": {},
	"none":    {},
	"pending": {},
	"n/a":     {},
	"na":      {},
	"tbd":     {},
	"-":       {},
}

func isFakePlaceholder(v string) bool {
	if v == "" {
		return false
	}
	_, ok := fakePlaceholders[strings.ToLower(strings.TrimSpace(v))]
	return ok
}
