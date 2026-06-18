package progressmessage

func messagesSorted(messages []MessageDefinition) bool {
	for i := 1; i < len(messages); i++ {
		if messages[i-1].Code > messages[i].Code {
			return false
		}
	}
	return true
}
