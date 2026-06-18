package progressmessage

import (
	"context"
	"os/exec"
	"strings"
	"time"
)

func baselineCatalogFromGit() ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if base, err := exec.CommandContext(ctx, "git", "merge-base", "HEAD", "origin/main").Output(); err == nil {
		ref := strings.TrimSpace(string(base))
		if ref != "" {
			out, err := exec.CommandContext(ctx, "git", "show", ref+":progressmessage/catalog.ir.riido.json").Output()
			if err == nil {
				return out, nil
			}
		}
	}
	return exec.CommandContext(ctx, "git", "show", "HEAD^:progressmessage/catalog.ir.riido.json").Output()
}
