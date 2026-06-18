package apicontract

func copyClientModules(modules []ClientModule) []ClientModule {
	if len(modules) == 0 {
		return nil
	}
	out := make([]ClientModule, 0, len(modules))
	for _, module := range modules {
		copied := ClientModule{
			Module:      module.Module,
			Description: module.Description,
			Namespaces:  make([]ClientNamespace, 0, len(module.Namespaces)),
		}
		for _, namespace := range module.Namespaces {
			copied.Namespaces = append(copied.Namespaces, ClientNamespace{
				Path:        append([]string(nil), namespace.Path...),
				Description: namespace.Description,
			})
		}
		out = append(out, copied)
	}
	return out
}
