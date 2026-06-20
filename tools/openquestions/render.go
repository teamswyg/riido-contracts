package main

import (
	"fmt"
	"strings"
)

func renderDoc(m manifest) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# %s\n\n%s\n\n", m.Title, generatedNotice)
	fmt.Fprintf(&b, "> Riido task: %s\n\n", m.RiidoTask)
	fmt.Fprintf(&b, "Executable SSOT: [`open-questions.riido.json`](open-questions.riido.json).\n\n")
	fmt.Fprintf(&b, "- Status: `%s`\n", status(m))
	fmt.Fprintf(&b, "- Open questions: `%d`\n", countStatus(m, "open"))
	fmt.Fprintf(&b, "- Resolved questions: `%d`\n\n", countStatus(m, "resolved"))
	renderQuestions(&b, m.Questions)
	renderLoop(&b, m.Loop)
	return b.String()
}

func renderQuestions(b *strings.Builder, questions []question) {
	b.WriteString("## Questions\n\n")
	b.WriteString("| ID | Status | Area | Question | Current stance | Next artifact |\n")
	b.WriteString("| --- | --- | --- | --- | --- | --- |\n")
	for _, q := range questions {
		fmt.Fprintf(b, "| `%s` | `%s` | %s | %s | %s | %s |\n",
			q.ID, q.Status, cell(q.Area), cell(q.Question), cell(q.CurrentStance), cell(q.NextArtifact))
	}
	b.WriteString("\n")
}
