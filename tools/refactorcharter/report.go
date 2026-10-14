package main

import (
	"fmt"
	"io"
)

const maxPrintedFindings = 20

func writeReport(out io.Writer, c charter, report scanReport) {
	fmt.Fprintf(
		out,
		"refactorcharter: mode=%s scanned=%d over_target=%d target_max_lines=%d\n",
		c.Mode,
		report.FilesScanned,
		len(report.Findings),
		c.LineBudget.TargetMaxLines,
	)
	limit := min(len(report.Findings), maxPrintedFindings)
	for i := 0; i < limit; i++ {
		f := report.Findings[i]
		fmt.Fprintf(out, "refactorcharter: finding lines=%d path=%s\n", f.Lines, f.Path)
	}
}

func enforced(c charter) bool {
	return c.Mode == "enforced"
}
