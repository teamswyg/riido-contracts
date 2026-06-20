package main

import "path/filepath"

func loadSingleIncludes(base string, c *contract) error {
	if c.ExecutionIdentityFile != "" {
		var doc executionIDDocument
		path := filepath.Join(base, c.ExecutionIdentityFile)
		if err := loadInclude(path, "execution identity", executionIDSchema, &doc); err != nil {
			return err
		}
		c.ExecutionIdentity = doc.ExecutionID
	}
	if c.ApprovalContractFile != "" {
		var doc approvalContractDocument
		path := filepath.Join(base, c.ApprovalContractFile)
		if err := loadInclude(path, "approval contract", approvalSchema, &doc); err != nil {
			return err
		}
		c.ApprovalContract = doc.Approval
	}
	return nil
}
