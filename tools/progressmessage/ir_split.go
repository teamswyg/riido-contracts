package main

import "github.com/teamswyg/riido-contracts/progressmessage"

type generatedIRFiles struct {
	Root     []byte
	Messages map[string][]byte
}

func generatedIRFilesFor(ir progressmessage.IRDocument) (generatedIRFiles, error) {
	root := ir
	root.MessageFiles = irMessageFileRefs(ir.Messages)
	root.Messages = nil
	body, err := progressmessage.MarshalCanonical(root)
	if err != nil {
		return generatedIRFiles{}, err
	}
	messages, err := generatedIRMessages(ir.Messages)
	if err != nil {
		return generatedIRFiles{}, err
	}
	return generatedIRFiles{Root: body, Messages: messages}, nil
}

func generatedIRMessages(messages []progressmessage.MessageDefinition) (map[string][]byte, error) {
	out := map[string][]byte{}
	for _, message := range messages {
		doc := progressmessage.MessageDocument{
			SchemaVersion: progressmessage.IRMessageSchemaVersion,
			Message:       message,
		}
		body, err := progressmessage.MarshalCanonical(doc)
		if err != nil {
			return nil, err
		}
		out[irMessageFile(message)] = body
	}
	return out, nil
}
