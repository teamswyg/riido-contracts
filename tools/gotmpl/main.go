package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"
)

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "gotmpl:", err)
		os.Exit(1)
	}
}

func run(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("gotmpl", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	bodyPath := fs.String("body", "", "template body filepath")
	dataPath := fs.String("data", "", "template data JSON filepath")
	outputPath := fs.String("out", "", "output filepath; stdout when empty")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *bodyPath == "" || *dataPath == "" {
		return errors.New("usage: go tool gotmpl -body TEMPLATE.gotmpl -data DATA.json [-out OUTPUT]")
	}
	body, err := os.ReadFile(*bodyPath)
	if err != nil {
		return fmt.Errorf("read body: %w", err)
	}
	dataBody, err := os.ReadFile(*dataPath)
	if err != nil {
		return fmt.Errorf("read data: %w", err)
	}
	var data any
	if err := json.Unmarshal(dataBody, &data); err != nil {
		return fmt.Errorf("decode data: %w", err)
	}
	tmpl, err := template.New(filepath.Base(*bodyPath)).Option("missingkey=error").Parse(string(body))
	if err != nil {
		return fmt.Errorf("parse template: %w", err)
	}
	target := out
	var file *os.File
	if *outputPath != "" {
		if err := os.MkdirAll(filepath.Dir(*outputPath), 0o755); err != nil {
			return fmt.Errorf("mkdir output: %w", err)
		}
		file, err = os.Create(*outputPath)
		if err != nil {
			return fmt.Errorf("create output: %w", err)
		}
		defer file.Close()
		target = file
	}
	if err := tmpl.Execute(target, data); err != nil {
		return fmt.Errorf("execute template: %w", err)
	}
	return nil
}
