package progressmessage

import (
	"fmt"
	"io/fs"
	"path"
)

func LoadDSL(fsys fs.FS, file string) (DSLDocument, error) {
	var dsl DSLDocument
	if err := loadCatalogJSON(fsys, file, &dsl); err != nil {
		return DSLDocument{}, err
	}
	if err := loadDSLMessageFiles(fsys, path.Dir(file), &dsl); err != nil {
		return DSLDocument{}, err
	}
	if err := ValidateDSL(dsl); err != nil {
		return DSLDocument{}, err
	}
	return dsl, nil
}

func LoadIR(fsys fs.FS, file string) (IRDocument, error) {
	var ir IRDocument
	if err := loadCatalogJSON(fsys, file, &ir); err != nil {
		return IRDocument{}, err
	}
	if err := loadIRMessageFiles(fsys, path.Dir(file), &ir); err != nil {
		return IRDocument{}, err
	}
	if err := ValidateIR(ir); err != nil {
		return IRDocument{}, err
	}
	return ir, nil
}

func loadCatalogJSON(fsys fs.FS, file string, dest any) error {
	body, err := fs.ReadFile(fsys, file)
	if err != nil {
		return fmt.Errorf("progressmessage: read %s: %w", file, err)
	}
	return decodeStrictJSON(file, body, dest)
}
