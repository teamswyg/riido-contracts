package main

import "strings"

func markdownCell(s string) string {
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "|", "\\|")
	return s
}

func renderOwner(o ownerRef) string {
	return "`" + markdownCell(o.Repo) + ":" + markdownCell(o.Path) + "`"
}

func renderSourceRef(ref sourceRef) string {
	return "`" + markdownCell(ref.Repo) + ":" + markdownCell(ref.Path) + "` requires `" +
		markdownCell(ref.RequiredPhrase) + "`"
}

func joinDownstreamRepos(downstreams []downstream) string {
	repos := make([]string, 0, len(downstreams))
	for _, downstream := range downstreams {
		repos = append(repos, downstream.Repo)
	}
	return strings.Join(repos, ", ")
}
