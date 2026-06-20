package main

import (
	"fmt"
	"io"
)

func runRender(args []string, out io.Writer) error {
	fs := quietFlagSet("render")
	manifestPath := manifestFlag(fs)
	if err := fs.Parse(args); err != nil {
		return err
	}
	_, m, err := loadDefaultedManifest(*manifestPath)
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(out, renderManifest(m))
	return err
}
