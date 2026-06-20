package main

import "errors"

func verifyContextEntries(m manifest) error {
	for _, ctx := range m.OwnedContexts {
		if blank(ctx.Context) || blank(ctx.Package) || blank(ctx.Responsibility) {
			return errors.New("owned context entries must be complete")
		}
	}
	for _, ctx := range m.NonOwnedContexts {
		if blank(ctx.Context) || blank(ctx.Owner) || blank(ctx.Boundary) {
			return errors.New("non-owned context entries must be complete")
		}
	}
	for _, link := range m.SSOTLinks {
		if blank(link.Label) || blank(link.Path) {
			return errors.New("ssot link entries must be complete")
		}
	}
	return nil
}
