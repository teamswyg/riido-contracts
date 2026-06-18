package apicontract

import (
	"strings"
)

func copyClientMeta(meta *ClientMeta) *ClientMeta {
	if meta == nil {
		return nil
	}
	out := *meta
	out.FacadePath = append([]string(nil), meta.FacadePath...)
	out.Invalidates = append([]string(nil), meta.Invalidates...)
	return &out
}

func deriveClientMeta(meta *ClientMeta) *ClientMeta {
	out := copyClientMeta(meta)
	if out == nil {
		return nil
	}
	out.GeneratedPath = generatedClientPath(*out)
	return out
}

func generatedClientPath(meta ClientMeta) string {
	if strings.TrimSpace(meta.Module) == "" || len(meta.FacadePath) == 0 {
		return ""
	}
	return meta.Module + "." + strings.Join(meta.FacadePath, ".")
}
