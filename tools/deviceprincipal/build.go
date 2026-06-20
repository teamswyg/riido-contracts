package main

import (
	"fmt"

	"github.com/teamswyg/riido-contracts/deviceprincipal"
)

func build(root string, opt options) (model, string, error) {
	m, err := readJSONFile[manifest](resolve(root, opt.manifest))
	if err != nil {
		return model{}, "", err
	}
	if err := verifyManifest(m); err != nil {
		return model{}, "", err
	}
	model, err := buildModel(root, m)
	if err != nil {
		return model, "", err
	}
	if err := verifyModel(model); err != nil {
		return model, "", err
	}
	return model, renderDoc(model), nil
}

func buildModel(root string, m manifest) (model, error) {
	rules, err := loadPolicyRules(root, m.APIPolicyID)
	if err != nil {
		return model{}, err
	}
	depCount, err := verifyDependencyFact(root, m.DependencyFactID)
	if err != nil {
		return model{}, err
	}
	return model{
		Manifest:              m,
		PrincipalKinds:        principalStrings(),
		DaemonHeaders:         deviceprincipal.DaemonCredentialHeaders(),
		ClientHeaders:         deviceprincipal.ClientCredentialHeaders(),
		OwnershipEdges:        ownershipStrings(),
		BindingSources:        deviceprincipal.BindingSources(),
		ExcludedFallbacks:     deviceprincipal.ExcludedFallbacks(),
		SecretSinks:           deviceprincipal.SecretNonExposureSinks(),
		DependencyPhrases:     deviceprincipal.DependencyPhrases(),
		BindingFields:         bindingJSONFields(),
		PolicyRules:           rules,
		SnapshotInterval:      deviceprincipal.RuntimeSnapshotIntervalSeconds,
		RuntimeStaleAfter:     deviceprincipal.RuntimeStaleAfterSeconds,
		DependencyPhraseCount: depCount,
	}, nil
}

func countMismatch(name string, got, want int) error {
	if got != want {
		return fmt.Errorf("%s = %d, want %d", name, got, want)
	}
	return nil
}
