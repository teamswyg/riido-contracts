package progressmessage

import (
	"context"
	"fmt"
	"os/exec"
	"path"
	"strings"
	"time"
)

func baselineCatalogFromGit() (IRDocument, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if base, err := exec.CommandContext(ctx, "git", "merge-base", "HEAD", "origin/main").Output(); err == nil {
		ref := strings.TrimSpace(string(base))
		if ref != "" {
			ir, err := catalogFromGitRef(ctx, ref)
			if err == nil {
				return ir, nil
			}
		}
	}
	return catalogFromGitRef(ctx, "HEAD^")
}

func catalogFromGitRef(ctx context.Context, ref string) (IRDocument, error) {
	ir, err := readGitIR(ctx, ref, "progressmessage/catalog.ir.riido.json")
	if err != nil {
		return IRDocument{}, err
	}
	for _, file := range ir.MessageFiles {
		doc, err := readGitMessage(ctx, ref, path.Join("progressmessage", file))
		if err != nil {
			return IRDocument{}, err
		}
		ir.Messages = append(ir.Messages, doc.Message)
	}
	return ir, ValidateIR(ir)
}

func readGitIR(ctx context.Context, ref, file string) (IRDocument, error) {
	body, err := exec.CommandContext(ctx, "git", "show", ref+":"+file).Output()
	if err != nil {
		return IRDocument{}, err
	}
	var ir IRDocument
	return ir, decodeStrictJSON(file, body, &ir)
}

func readGitMessage(ctx context.Context, ref, file string) (MessageDocument, error) {
	body, err := exec.CommandContext(ctx, "git", "show", ref+":"+file).Output()
	if err != nil {
		return MessageDocument{}, err
	}
	var doc MessageDocument
	if err := decodeStrictJSON(file, body, &doc); err != nil {
		return MessageDocument{}, err
	}
	if doc.SchemaVersion != IRMessageSchemaVersion {
		return MessageDocument{}, fmt.Errorf("%s schema_version = %q", file, doc.SchemaVersion)
	}
	return doc, nil
}
