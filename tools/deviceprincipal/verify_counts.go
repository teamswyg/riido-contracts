package main

func verifyCounts(model model) error {
	m := model.Manifest
	checks := map[string][2]int{
		"principal count":          {len(model.PrincipalKinds), m.ExpectedPrincipalCount},
		"daemon headers":           {len(model.DaemonHeaders), m.ExpectedDaemonHeaderCount},
		"client headers":           {len(model.ClientHeaders), m.ExpectedClientHeaderCount},
		"snapshot interval":        {model.SnapshotInterval, m.ExpectedSnapshotIntervalSeconds},
		"runtime stale after":      {model.RuntimeStaleAfter, m.ExpectedRuntimeStaleAfterSeconds},
		"ownership edges":          {len(model.OwnershipEdges), m.ExpectedOwnershipEdgeCount},
		"binding sources":          {len(model.BindingSources), m.ExpectedBindingSourceCount},
		"excluded fallbacks":       {len(model.ExcludedFallbacks), m.ExpectedExcludedFallbackCount},
		"secret sinks":             {len(model.SecretSinks), m.ExpectedSecretNonExposureSinks},
		"dependency phrases":       {model.DependencyPhraseCount, m.ExpectedDependencyPhraseCount},
		"runtime binding fields":   {len(model.BindingFields), m.ExpectedBindingFieldCount},
		"policy rule prefix count": {len(m.PolicyRulePrefixes), 8},
	}
	for name, values := range checks {
		if err := countMismatch(name, values[0], values[1]); err != nil {
			return err
		}
	}
	return nil
}
