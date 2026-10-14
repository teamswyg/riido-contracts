package main

import (
	"fmt"
	"strings"
)

func verifyRepoDependencies(deps []repoDependency, factIDs map[string]bool) error {
	seen := map[string]bool{}
	graph := map[string][]string{}
	for _, dep := range deps {
		if err := verifyRepoDependency(dep, factIDs, seen); err != nil {
			return err
		}
		graph[dep.FromRepo] = append(graph[dep.FromRepo], dep.ToRepo)
		if _, ok := graph[dep.ToRepo]; !ok {
			graph[dep.ToRepo] = nil
		}
	}
	return verifyAcyclic(graph)
}

func verifyRepoDependency(dep repoDependency, factIDs, seen map[string]bool) error {
	if err := requireID("repo_dependency id", dep.ID); err != nil {
		return err
	}
	if seen[dep.ID] {
		return fmt.Errorf("duplicate repo_dependency id %q", dep.ID)
	}
	seen[dep.ID] = true
	if strings.TrimSpace(dep.FromRepo) == "" || strings.TrimSpace(dep.ToRepo) == "" {
		return fmt.Errorf("repo_dependency %q from_repo and to_repo are required", dep.ID)
	}
	if dep.FromRepo == dep.ToRepo {
		return fmt.Errorf("repo_dependency %q cannot point to itself", dep.ID)
	}
	if strings.TrimSpace(dep.LocalScope) == "" {
		return fmt.Errorf("repo_dependency %q local_scope is required", dep.ID)
	}
	if len(dep.FactIDs) == 0 {
		return fmt.Errorf("repo_dependency %q fact_ids are required", dep.ID)
	}
	if !stringsSorted(dep.FactIDs) {
		return fmt.Errorf("repo_dependency %q fact_ids must be sorted", dep.ID)
	}
	for _, factID := range dep.FactIDs {
		if !factIDs[factID] {
			return fmt.Errorf("repo_dependency %q references unknown fact %q", dep.ID, factID)
		}
	}
	return nil
}
