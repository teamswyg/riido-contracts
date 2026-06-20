package main

import "fmt"

func verifyFixtureRows(rows []fixtureRow) error {
	for _, row := range rows {
		if row.Name == "" || row.TmpColor == "" {
			return fmt.Errorf("fixture rows require name and tmp_color")
		}
		if row.Name == "직접 설정" {
			return fmt.Errorf("direct setting must not be an onboarding fixture")
		}
	}
	return nil
}
