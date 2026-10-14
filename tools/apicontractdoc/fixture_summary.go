package main

import (
	"strings"

	"github.com/teamswyg/riido-contracts/apicontract"
)

func summarizeFixture(ref fixtureRef, ir apicontract.IRDocument) fixtureSummary {
	summary := fixtureSummary{
		ContractID:     ref.ContractID,
		Context:        ir.Context,
		ServiceSchema:  ir.Service.SchemaVersion,
		OperationCount: len(ir.Operations),
		ComponentCount: len(ir.Components),
		EnumCount:      len(ir.Enums),
		SumTypeCount:   len(ir.SumTypes),
	}
	for _, operation := range ir.Operations {
		if operation.Client != nil && operation.Client.GeneratedPath != "" {
			summary.GeneratedPathCount++
		}
		if strings.HasPrefix(operation.Path, "/v2/") {
			summary.V2OperationCount++
		}
	}
	return summary
}

func collectGeneratedPaths(out map[string]bool, ir apicontract.IRDocument) {
	for _, operation := range ir.Operations {
		if operation.Client != nil && operation.Client.GeneratedPath != "" {
			out[operation.Client.GeneratedPath] = true
		}
	}
}
