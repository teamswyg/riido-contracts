package main

func specHasMultiTargetEvent(spec fsmMetadata) bool {
	if spec.EventEnum == "" {
		return false
	}
	targetsByFromEvent := map[string]map[string]bool{}
	for _, entry := range spec.Entries {
		key := entry.From + "\x00" + entry.Event
		if _, ok := targetsByFromEvent[key]; !ok {
			targetsByFromEvent[key] = map[string]bool{}
		}
		targetsByFromEvent[key][entry.To] = true
		if len(targetsByFromEvent[key]) > 1 {
			return true
		}
	}
	return false
}

func reachableVertices(starts []string, outgoing map[string][]string) map[string]bool {
	reachable := map[string]bool{}
	queue := append([]string(nil), starts...)
	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]
		if reachable[state] {
			continue
		}
		reachable[state] = true
		queue = append(queue, outgoing[state]...)
	}
	return reachable
}

func stringSet(values []string) map[string]bool {
	out := make(map[string]bool, len(values))
	for _, value := range values {
		out[value] = true
	}
	return out
}
