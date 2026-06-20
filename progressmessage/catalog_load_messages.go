package progressmessage

import (
	"fmt"
	"io/fs"
	"path"
)

func loadDSLMessageFiles(fsys fs.FS, base string, dsl *DSLDocument) error {
	for _, file := range dsl.MessageFiles {
		doc, err := loadMessageDocument(fsys, base, file, DSLMessageSchemaVersion)
		if err != nil {
			return err
		}
		dsl.Messages = append(dsl.Messages, doc.Message)
	}
	return nil
}

func loadIRMessageFiles(fsys fs.FS, base string, ir *IRDocument) error {
	for _, file := range ir.MessageFiles {
		doc, err := loadMessageDocument(fsys, base, file, IRMessageSchemaVersion)
		if err != nil {
			return err
		}
		ir.Messages = append(ir.Messages, doc.Message)
	}
	return nil
}

func loadMessageDocument(
	fsys fs.FS,
	base string,
	file string,
	version string,
) (MessageDocument, error) {
	var doc MessageDocument
	file = path.Join(base, file)
	if err := loadCatalogJSON(fsys, file, &doc); err != nil {
		return MessageDocument{}, err
	}
	if doc.SchemaVersion != version {
		return MessageDocument{}, fmt.Errorf("%s schema_version = %q", file, doc.SchemaVersion)
	}
	return doc, nil
}
