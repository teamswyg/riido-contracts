package main

import (
	"fmt"
	"strings"
)

type readmeSectionMarkerBounds struct {
	ContentStart int
	ContentEnd   int
}

func readmeSectionBounds(body, start, end string) (readmeSectionMarkerBounds, error) {
	startIndex := strings.Index(body, start)
	if startIndex < 0 {
		return readmeSectionMarkerBounds{}, fmt.Errorf("%s missing marker %s", readmePath, start)
	}
	contentStart := startIndex + len(start)
	endIndex := strings.Index(body[contentStart:], end)
	if endIndex < 0 {
		return readmeSectionMarkerBounds{}, fmt.Errorf("%s missing marker %s", readmePath, end)
	}
	return readmeSectionMarkerBounds{
		ContentStart: contentStart,
		ContentEnd:   contentStart + endIndex,
	}, nil
}
