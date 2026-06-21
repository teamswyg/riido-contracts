package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

func run(args []string) error {
	fs := flag.NewFlagSet("cievidence", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	workflow := fs.String("workflow", "", "workflow path")
	id := fs.String("id", "", "evidence id")
	evidenceOut := fs.String("evidence-out", "", "evidence output path")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *workflow == "" || *id == "" || *evidenceOut == "" {
		return errors.New("-workflow, -id, and -evidence-out are required")
	}
	body, err := os.ReadFile(*workflow)
	if err != nil {
		return fmt.Errorf("read workflow: %w", err)
	}
	value := buildEvidence(*id, *workflow, string(body))
	if err := writeEvidence(*evidenceOut, value); err != nil {
		return err
	}
	if value.Status != "verified" {
		return fmt.Errorf("%s evidence status %s", *id, value.Status)
	}
	return nil
}
