package main

import (
	"fmt"
	"strings"
)

func renderFSM(b *strings.Builder, fsm fsmDoc) {
	b.WriteString("## Generated FSM\n\n")
	renderParagraphs(b, fsm.Intro)
	for _, section := range fsm.Sections {
		fmt.Fprintf(b, "%s:\n\n%s\n\n", section.Title, section.Body)
	}
}

func renderCommands(b *strings.Builder, commands []string) {
	b.WriteString("## 검증\n\n```bash\n")
	for _, command := range commands {
		b.WriteString(command + "\n")
	}
	b.WriteString("```\n\n")
}

func renderLoop(b *strings.Builder, loop evidenceLoop) {
	b.WriteString("## Evidence Loop\n\n| Step | Statement |\n| --- | --- |\n")
	fmt.Fprintf(b, "| Observe | %s |\n", loop.Observation)
	fmt.Fprintf(b, "| Hypothesis | %s |\n", loop.Hypothesis)
	fmt.Fprintf(b, "| Execute | %s |\n", loop.Execute)
	fmt.Fprintf(b, "| Evaluate | %s |\n", loop.Evaluate)
	fmt.Fprintf(b, "| Retrospective | %s |\n\n", loop.Retrospective)
}
