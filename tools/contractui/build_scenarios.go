package main

import "strings"

func attachScenarios(ops []operation, dsl dslDocument) {
	scenarios := map[string][]scenario{}
	for _, op := range dsl.Operations {
		scenarios[op.OperationID] = op.Scenarios
	}
	for i := range ops {
		ops[i].Scenarios = scenarios[ops[i].OperationID]
	}
}

func dslPathForIR(irPath string) string {
	return strings.TrimSuffix(irPath, ".ir.riido.json") + ".dsl.riido.json"
}
