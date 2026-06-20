package main

import "github.com/teamswyg/riido-contracts/progressmessage"

func status(ir progressmessage.IRDocument) string {
	if len(ir.Messages) < ir.MaxMessages {
		return "verified_with_capacity"
	}
	return "verified"
}

func usageCounts(ir progressmessage.IRDocument) map[string]int {
	counts := map[string]int{}
	for _, message := range ir.Messages {
		counts[message.Usage]++
	}
	return counts
}
