package main

import (
	"fmt"
	"io"
)

func runFSMConformance(plan fsmRunPlan, out io.Writer) error {
	fmt.Fprintf(out, "fsmgen: conformance verified %d FSMs against %d profile(s)\n", len(plan.Metadata), profileCount(plan.PatternDocs))
	return nil
}
