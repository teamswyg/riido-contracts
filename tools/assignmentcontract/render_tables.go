package main

import "strings"

func renderContractTables(b *strings.Builder, c contract) {
	renderStates(b, c.AssignmentStates)
	renderNamedValues(b, "Poll Actions", c.PollActions)
	renderPayloadFields(b, c.AssignmentPayloadFields)
}
