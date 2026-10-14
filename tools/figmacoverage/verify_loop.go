package main

func completeLoop(loop evidenceLoop) bool {
	return !blank(loop.Observation) &&
		!blank(loop.Hypothesis) &&
		!blank(loop.Execute) &&
		!blank(loop.Evaluate) &&
		!blank(loop.Retrospective)
}
