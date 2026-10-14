package main

import "errors"

func verifyRenderableManifest(m manifest) error {
	if m.SchemaVersion != schemaVersion {
		return errors.New("unsupported schema_version")
	}
	if err := requireID("manifest id", m.ID); err != nil {
		return err
	}
	if len(m.Facts) == 0 {
		return errors.New("facts are required")
	}
	return nil
}
