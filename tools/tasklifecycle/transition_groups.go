package main

import (
	"strings"

	"github.com/teamswyg/riido-contracts/task"
)

func groupTransitions(transitions []task.TaskTransitionCode) []transitionGroup {
	var groups []transitionGroup
	groupByFrom := map[string]int{}
	edgeByKey := map[string]int{}
	for _, transition := range transitions {
		from, trigger := transition.From.String(), transition.Trigger.String()
		gi := ensureGroup(&groups, groupByFrom, from)
		key := from + "\x00" + trigger
		ei, ok := edgeByKey[key]
		if !ok {
			ei = len(groups[gi].Edges)
			edgeByKey[key] = ei
			groups[gi].Edges = append(groups[gi].Edges, transitionEdge{Trigger: trigger})
		}
		groups[gi].Edges[ei].To = append(groups[gi].Edges[ei].To, transition.To.String())
	}
	return groups
}

func ensureGroup(groups *[]transitionGroup, index map[string]int, from string) int {
	if i, ok := index[from]; ok {
		return i
	}
	i := len(*groups)
	index[from] = i
	*groups = append(*groups, transitionGroup{From: from})
	return i
}

func renderEdge(edge transitionEdge) string {
	return edge.Trigger + " -> " + strings.Join(edge.To, "/")
}
