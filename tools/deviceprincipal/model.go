package main

type model struct {
	Manifest              manifest
	PrincipalKinds        []string
	DaemonHeaders         []string
	ClientHeaders         []string
	OwnershipEdges        []string
	BindingSources        []string
	ExcludedFallbacks     []string
	SecretSinks           []string
	DependencyPhrases     []string
	BindingFields         []string
	PolicyRules           []string
	SnapshotInterval      int
	RuntimeStaleAfter     int
	DependencyPhraseCount int
}
