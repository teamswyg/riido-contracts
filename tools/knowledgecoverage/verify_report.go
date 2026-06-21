package main

import "fmt"

func verifyReport(report scanReport) error {
	if report.ManualCount == 0 {
		return nil
	}
	return fmt.Errorf("manual reader candidates found: %d", report.ManualCount)
}
