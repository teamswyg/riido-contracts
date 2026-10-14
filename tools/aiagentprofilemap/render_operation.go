package main

import (
	"fmt"
	"strings"
)

func renderOperation(b *strings.Builder, op operation) {
	b.WriteString("## Operation\n\n")
	fmt.Fprintf(b, "- `%s` `%s %s` -> `%d %s`; generated path `%s`; scenarios `%d`\n\n",
		op.OperationID, op.Method, op.Path, responseStatus(op), responseName(op),
		clientRoute(op), len(op.Scenarios))
	renderScenarioNames(b, op.Scenarios)
}

func renderScenarioNames(b *strings.Builder, scenarios []scenario) {
	names := make([]string, 0, len(scenarios))
	for _, scenario := range scenarios {
		names = append(names, scenario.Name)
	}
	fmt.Fprintf(b, "- Scenario names: `%s`\n\n", strings.Join(names, "`, `"))
}
