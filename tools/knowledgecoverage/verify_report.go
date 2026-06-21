package main

import "fmt"

func verifyReport(report scanReport) error {
	if len(report.Problems) > 0 {
		return fmt.Errorf("%s", report.Problems[0])
	}
	if report.ManualCount == 0 {
		return nil
	}
	return fmt.Errorf("manual reader candidates found: %d", report.ManualCount)
}
