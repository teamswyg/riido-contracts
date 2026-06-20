package main

import (
	"fmt"

	"github.com/teamswyg/riido-contracts/deviceprincipal"
)

func principalStrings() []string {
	out := make([]string, 0, len(deviceprincipal.PrincipalKinds()))
	for _, kind := range deviceprincipal.PrincipalKinds() {
		out = append(out, string(kind))
	}
	return out
}

func ownershipStrings() []string {
	out := make([]string, 0, len(deviceprincipal.OwnershipEdges()))
	for _, edge := range deviceprincipal.OwnershipEdges() {
		out = append(out, fmt.Sprintf("%s -> %s", edge.From, edge.To))
	}
	return out
}
