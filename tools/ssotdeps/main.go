package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

const (
	schemaVersion   = "riido-ssot-dependency-map.v1"
	defaultManifest = "docs/30-architecture/ssot-dependency-map.riido.json"
	localRepo       = "riido-contracts"
)

var idPattern = regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)

type manifest struct {
	SchemaVersion    string           `json:"schema_version"`
	ID               string           `json:"id"`
	RiidoTask        string           `json:"riido_task"`
	HumanDoc         string           `json:"human_doc"`
	Facts            []fact           `json:"facts"`
	RepoDependencies []repoDependency `json:"repo_dependencies"`
}

type fact struct {
	ID             string       `json:"id"`
	Fact           string       `json:"fact"`
	HumanDocPhrase string       `json:"human_doc_phrase"`
	SourceRefs     []sourceRef  `json:"source_refs"`
	Owner          ownerRef     `json:"owner"`
	Downstreams    []downstream `json:"downstreams"`
}

type sourceRef struct {
	Repo           string `json:"repo"`
	Path           string `json:"path"`
	RequiredPhrase string `json:"required_phrase"`
}

type ownerRef struct {
	Repo string `json:"repo"`
	Path string `json:"path"`
}

type downstream struct {
	Repo       string `json:"repo"`
	LocalScope string `json:"local_scope"`
}

type repoDependency struct {
	ID         string   `json:"id"`
	FromRepo   string   `json:"from_repo"`
	ToRepo     string   `json:"to_repo"`
	FactIDs    []string `json:"fact_ids"`
	LocalScope string   `json:"local_scope"`
}

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "ssotdeps:", err)
		os.Exit(1)
	}
}

func run(args []string, out io.Writer) error {
	command := "verify"
	if len(args) > 0 && !strings.HasPrefix(args[0], "-") {
		command = args[0]
		args = args[1:]
	}
	switch command {
	case "verify":
		return runVerify(args, out)
	default:
		return fmt.Errorf("usage: go run ./tools/ssotdeps [verify] [-manifest %s]", defaultManifest)
	}
}

func runVerify(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("verify", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	manifestPath := fs.String("manifest", defaultManifest, "SSOT dependency manifest path")
	if err := fs.Parse(args); err != nil {
		return err
	}
	root := "."
	loadPath := *manifestPath
	if *manifestPath == defaultManifest {
		repoRoot, err := findRepoRoot()
		if err != nil {
			return err
		}
		root = repoRoot
		loadPath = filepath.Join(repoRoot, filepath.FromSlash(*manifestPath))
	}
	m, err := loadManifest(loadPath)
	if err != nil {
		return err
	}
	if err := verifyManifest(m, root); err != nil {
		return err
	}
	fmt.Fprintf(out, "ssotdeps: verified %d facts and %d repo dependencies\n", len(m.Facts), len(m.RepoDependencies))
	return nil
}

func loadManifest(path string) (manifest, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return manifest{}, fmt.Errorf("read manifest: %w", err)
	}
	dec := json.NewDecoder(strings.NewReader(string(body)))
	dec.DisallowUnknownFields()
	var m manifest
	if err := dec.Decode(&m); err != nil {
		return manifest{}, fmt.Errorf("decode manifest: %w", err)
	}
	var extra struct{}
	if err := dec.Decode(&extra); err != io.EOF {
		return manifest{}, errors.New("decode manifest: trailing data")
	}
	return m, nil
}

func verifyManifest(m manifest, root string) error {
	if m.SchemaVersion != schemaVersion {
		return fmt.Errorf("schema_version = %q, want %q", m.SchemaVersion, schemaVersion)
	}
	if err := requireID("manifest id", m.ID); err != nil {
		return err
	}
	if strings.TrimSpace(m.RiidoTask) == "" {
		return errors.New("riido_task is required")
	}
	humanDoc, err := readLocalRef(root, m.HumanDoc)
	if err != nil {
		return fmt.Errorf("human_doc: %w", err)
	}
	if len(m.Facts) == 0 {
		return errors.New("facts are required")
	}
	if !factsSorted(m.Facts) {
		return errors.New("facts must be sorted by id")
	}
	factIDs := map[string]bool{}
	for _, f := range m.Facts {
		if err := verifyFact(root, humanDoc, f); err != nil {
			return fmt.Errorf("fact %q: %w", f.ID, err)
		}
		if factIDs[f.ID] {
			return fmt.Errorf("duplicate fact id %q", f.ID)
		}
		factIDs[f.ID] = true
	}
	if !repoDependenciesSorted(m.RepoDependencies) {
		return errors.New("repo_dependencies must be sorted by id")
	}
	if err := verifyRepoDependencies(m.RepoDependencies, factIDs); err != nil {
		return err
	}
	return nil
}

func verifyFact(root, humanDoc string, f fact) error {
	if err := requireID("id", f.ID); err != nil {
		return err
	}
	if strings.TrimSpace(f.Fact) == "" {
		return errors.New("fact is required")
	}
	if strings.TrimSpace(f.HumanDocPhrase) == "" {
		return errors.New("human_doc_phrase is required")
	}
	if !strings.Contains(humanDoc, f.HumanDocPhrase) {
		return fmt.Errorf("human_doc does not contain phrase %q", f.HumanDocPhrase)
	}
	if err := verifyOwner(root, f.Owner); err != nil {
		return err
	}
	if len(f.SourceRefs) == 0 {
		return errors.New("source_refs are required")
	}
	seenRefs := map[string]bool{}
	for _, ref := range f.SourceRefs {
		key := ref.Repo + ":" + ref.Path + ":" + ref.RequiredPhrase
		if seenRefs[key] {
			return fmt.Errorf("duplicate source_ref %q", key)
		}
		seenRefs[key] = true
		if err := verifySourceRef(root, ref); err != nil {
			return err
		}
	}
	if len(f.Downstreams) == 0 {
		return errors.New("downstreams are required")
	}
	seenDownstreams := map[string]bool{}
	for _, downstream := range f.Downstreams {
		if strings.TrimSpace(downstream.Repo) == "" {
			return errors.New("downstream repo is required")
		}
		if strings.TrimSpace(downstream.LocalScope) == "" {
			return fmt.Errorf("downstream %q local_scope is required", downstream.Repo)
		}
		if seenDownstreams[downstream.Repo] {
			return fmt.Errorf("duplicate downstream repo %q", downstream.Repo)
		}
		seenDownstreams[downstream.Repo] = true
	}
	return nil
}

func verifyOwner(root string, owner ownerRef) error {
	if strings.TrimSpace(owner.Repo) == "" {
		return errors.New("owner.repo is required")
	}
	if strings.TrimSpace(owner.Path) == "" {
		return errors.New("owner.path is required")
	}
	if owner.Repo == localRepo {
		if _, err := readLocalRef(root, owner.Path); err != nil {
			return fmt.Errorf("owner %s: %w", owner.Path, err)
		}
	}
	return nil
}

func verifySourceRef(root string, ref sourceRef) error {
	if strings.TrimSpace(ref.Repo) == "" {
		return errors.New("source_ref.repo is required")
	}
	if strings.TrimSpace(ref.Path) == "" {
		return errors.New("source_ref.path is required")
	}
	if strings.TrimSpace(ref.RequiredPhrase) == "" {
		return fmt.Errorf("source_ref %s:%s required_phrase is required", ref.Repo, ref.Path)
	}
	if ref.Repo != localRepo {
		return nil
	}
	body, err := readLocalRef(root, ref.Path)
	if err != nil {
		return fmt.Errorf("source_ref %s: %w", ref.Path, err)
	}
	if !strings.Contains(body, ref.RequiredPhrase) {
		return fmt.Errorf("source_ref %s does not contain phrase %q", ref.Path, ref.RequiredPhrase)
	}
	return nil
}

func verifyRepoDependencies(deps []repoDependency, factIDs map[string]bool) error {
	seen := map[string]bool{}
	graph := map[string][]string{}
	for _, dep := range deps {
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
		graph[dep.FromRepo] = append(graph[dep.FromRepo], dep.ToRepo)
		if _, ok := graph[dep.ToRepo]; !ok {
			graph[dep.ToRepo] = nil
		}
	}
	return verifyAcyclic(graph)
}

func verifyAcyclic(graph map[string][]string) error {
	const (
		unseen = 0
		active = 1
		done   = 2
	)
	state := map[string]int{}
	var visit func(string) error
	visit = func(node string) error {
		switch state[node] {
		case active:
			return fmt.Errorf("repo dependency cycle detected at %s", node)
		case done:
			return nil
		}
		state[node] = active
		next := append([]string(nil), graph[node]...)
		sort.Strings(next)
		for _, child := range next {
			if err := visit(child); err != nil {
				return err
			}
		}
		state[node] = done
		return nil
	}
	nodes := make([]string, 0, len(graph))
	for node := range graph {
		nodes = append(nodes, node)
	}
	sort.Strings(nodes)
	for _, node := range nodes {
		if state[node] == unseen {
			if err := visit(node); err != nil {
				return err
			}
		}
	}
	return nil
}

func readLocalRef(root, path string) (string, error) {
	if filepath.IsAbs(path) || strings.Contains(path, "..") {
		return "", fmt.Errorf("path %q must be repo-relative", path)
	}
	body, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(path)))
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func findRepoRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			if _, err := os.Stat(filepath.Join(dir, filepath.FromSlash(defaultManifest))); err == nil {
				return dir, nil
			}
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", errors.New("could not find repository root")
		}
		dir = parent
	}
}

func requireID(label, id string) error {
	if !idPattern.MatchString(id) {
		return fmt.Errorf("%s must match %s", label, idPattern.String())
	}
	return nil
}

func factsSorted(facts []fact) bool {
	ids := make([]string, 0, len(facts))
	for _, fact := range facts {
		ids = append(ids, fact.ID)
	}
	return stringsSorted(ids)
}

func repoDependenciesSorted(deps []repoDependency) bool {
	ids := make([]string, 0, len(deps))
	for _, dep := range deps {
		ids = append(ids, dep.ID)
	}
	return stringsSorted(ids)
}

func stringsSorted(values []string) bool {
	return sort.SliceIsSorted(values, func(i, j int) bool {
		return values[i] < values[j]
	})
}
