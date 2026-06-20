package main

import (
	"flag"
	"fmt"
	"io"
)

func runRender(args []string, out io.Writer) error {
	fs := quietFlagSet("render")
	manifestPath := manifestFlag(fs)
	doc := fs.String("doc", "module", "doc to render: module or integration")
	if err := fs.Parse(args); err != nil {
		return err
	}
	_, m, err := loadDefaultedManifest(*manifestPath)
	if err != nil {
		return err
	}
	switch *doc {
	case "module":
		_, err = fmt.Fprint(out, renderModuleDoc(m))
	case "integration":
		_, err = fmt.Fprint(out, renderIntegrationDoc(m))
	default:
		err = flag.ErrHelp
	}
	return err
}
