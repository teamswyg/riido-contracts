package main

import "github.com/teamswyg/riido-contracts/progressmessage"

func newEvidence(m docManifest, ir progressmessage.IRDocument) evidence {
	return evidence{
		SchemaVersion: evidenceVersion, ID: m.ID, Status: status(ir),
		ContractID: ir.ContractID, MessageCount: len(ir.Messages),
		MaxMessages: ir.MaxMessages, UsageCounts: usageCounts(ir),
		GeneratedDoc: m.GeneratedDoc, EvidenceArtifact: m.EvidenceArtifact,
		Workflow: m.Workflow, DSL: m.DSL, IR: m.IR, Loop: m.Loop,
		Messages: evidenceMessages(ir.Messages),
	}
}

func evidenceMessages(messages []progressmessage.MessageDefinition) []evidenceMessage {
	out := make([]evidenceMessage, 0, len(messages))
	for _, message := range messages {
		out = append(out, evidenceMessage{
			Code: message.Code, Key: message.Key,
			Usage: message.Usage, Category: message.Category,
		})
	}
	return out
}
