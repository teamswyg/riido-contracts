package discovery

import (
	"encoding/base64"
	"errors"
	"fmt"
	"path"
	"strings"
	"time"

	providercatalog "github.com/teamswyg/riido-contracts/provider/catalog"
)

const (
	PolicySchemaVersion   = "riido-provider-discovery-policy.v1"
	EnvelopeSchemaVersion = "riido-signed-provider-discovery-policy.v1"
	SignatureAlgorithm    = "Ed25519"
	MaxPolicyLifetime     = 7 * 24 * time.Hour
)

type Root string

const (
	RootProcessPath        Root = "process-path"
	RootLoginShellPath     Root = "login-shell-path"
	RootSystemApplications Root = "system-applications"
	RootUserApplications   Root = "user-applications"
	RootUserHome           Root = "user-home"
	RootLocalAppData       Root = "local-app-data"
	RootAppData            Root = "app-data"
	RootProgramFiles       Root = "program-files"
	RootProgramFilesX86    Root = "program-files-x86"
)

type Policy struct {
	SchemaVersion string    `json:"schema_version"`
	Revision      string    `json:"revision"`
	IssuedAt      time.Time `json:"issued_at"`
	ExpiresAt     time.Time `json:"expires_at"`
	Rules         []Rule    `json:"rules"`
}

type Rule struct {
	Provider   providercatalog.Kind `json:"provider"`
	OS         string               `json:"os"`
	Arch       string               `json:"arch"`
	Candidates []Candidate          `json:"candidates"`
}

type Candidate struct {
	Root         Root   `json:"root"`
	RelativePath string `json:"relative_path"`
}

type SignedEnvelope struct {
	SchemaVersion string `json:"schema_version"`
	KeyID         string `json:"key_id"`
	Algorithm     string `json:"algorithm"`
	Payload       string `json:"payload"`
	Signature     string `json:"signature"`
}

func (p Policy) Validate() error {
	if p.SchemaVersion != PolicySchemaVersion {
		return fmt.Errorf("provider discovery: unsupported schema %q", p.SchemaVersion)
	}
	if !boundedID(p.Revision, 128) {
		return errors.New("provider discovery: invalid revision")
	}
	if p.IssuedAt.IsZero() || !p.ExpiresAt.After(p.IssuedAt) || p.ExpiresAt.Sub(p.IssuedAt) > MaxPolicyLifetime {
		return errors.New("provider discovery: invalid policy lifetime")
	}
	if len(p.Rules) == 0 || len(p.Rules) > 64 {
		return errors.New("provider discovery: rule count must be between 1 and 64")
	}
	seen := map[string]struct{}{}
	for _, rule := range p.Rules {
		if err := rule.validate(); err != nil {
			return err
		}
		key := string(rule.Provider) + "/" + rule.OS + "/" + rule.Arch
		if _, ok := seen[key]; ok {
			return fmt.Errorf("provider discovery: duplicate rule %s", key)
		}
		seen[key] = struct{}{}
	}
	return nil
}

func (p Policy) ValidateAt(now time.Time) error {
	if err := p.Validate(); err != nil {
		return err
	}
	now = now.UTC()
	if p.IssuedAt.After(now.Add(5 * time.Minute)) {
		return errors.New("provider discovery: policy issued in the future")
	}
	if !now.Before(p.ExpiresAt) {
		return errors.New("provider discovery: policy expired")
	}
	return nil
}

func (e SignedEnvelope) Validate() error {
	if e.SchemaVersion != EnvelopeSchemaVersion || e.Algorithm != SignatureAlgorithm {
		return errors.New("provider discovery: unsupported signed envelope")
	}
	if !boundedID(e.KeyID, 64) {
		return errors.New("provider discovery: invalid key id")
	}
	payload, err := base64.StdEncoding.DecodeString(e.Payload)
	if err != nil || len(payload) == 0 || len(payload) > 64<<10 {
		return errors.New("provider discovery: invalid payload")
	}
	signature, err := base64.StdEncoding.DecodeString(e.Signature)
	if err != nil || len(signature) != 64 {
		return errors.New("provider discovery: invalid signature")
	}
	return nil
}

func (r Rule) validate() error {
	switch r.Provider {
	case providercatalog.KindClaude, providercatalog.KindCodex, providercatalog.KindCursor, providercatalog.KindOpenClaw:
	default:
		return fmt.Errorf("provider discovery: unsupported provider %q", r.Provider)
	}
	if r.OS != "darwin" && r.OS != "windows" && r.OS != "linux" {
		return fmt.Errorf("provider discovery: unsupported os %q", r.OS)
	}
	if r.Arch != "amd64" && r.Arch != "arm64" {
		return fmt.Errorf("provider discovery: unsupported arch %q", r.Arch)
	}
	if len(r.Candidates) == 0 || len(r.Candidates) > 16 {
		return errors.New("provider discovery: candidate count must be between 1 and 16")
	}
	seen := map[string]struct{}{}
	for _, candidate := range r.Candidates {
		if err := candidate.validate(); err != nil {
			return err
		}
		if !candidate.Root.ValidForOS(r.OS) {
			return fmt.Errorf("provider discovery: root %q is invalid for os %q", candidate.Root, r.OS)
		}
		if !providerExecutableName(r.Provider, candidate.RelativePath) {
			return fmt.Errorf("provider discovery: executable name is invalid for provider %q", r.Provider)
		}
		key := string(candidate.Root) + "/" + candidate.RelativePath
		if _, ok := seen[key]; ok {
			return fmt.Errorf("provider discovery: duplicate candidate %s", key)
		}
		seen[key] = struct{}{}
	}
	return nil
}

func (c Candidate) validate() error {
	if !c.Root.Valid() {
		return fmt.Errorf("provider discovery: unsupported root %q", c.Root)
	}
	value := strings.TrimSpace(c.RelativePath)
	if value == "" || len(value) > 256 || strings.ContainsRune(value, 0) || portableAbsolute(value) {
		return errors.New("provider discovery: relative path is unsafe")
	}
	for _, r := range value {
		if r < 0x20 || r == 0x7f || r == ':' {
			return errors.New("provider discovery: relative path is unsafe")
		}
	}
	normalized := strings.ReplaceAll(value, `\`, "/")
	if path.Clean(normalized) != normalized || normalized == "." || strings.HasPrefix(normalized, "../") {
		return errors.New("provider discovery: relative path is unsafe")
	}
	return nil
}

func providerExecutableName(provider providercatalog.Kind, relativePath string) bool {
	name := path.Base(strings.ReplaceAll(relativePath, `\`, "/"))
	for _, suffix := range []string{"", ".exe", ".cmd", ".bat"} {
		if (provider == providercatalog.KindClaude && name == "claude"+suffix) ||
			(provider == providercatalog.KindCodex && name == "codex"+suffix) ||
			(provider == providercatalog.KindOpenClaw && name == "openclaw"+suffix) ||
			(provider == providercatalog.KindCursor && name == "cursor-agent"+suffix) {
			return true
		}
	}
	return false
}

func (r Root) Valid() bool {
	switch r {
	case RootProcessPath, RootLoginShellPath, RootSystemApplications, RootUserApplications,
		RootUserHome, RootLocalAppData, RootAppData, RootProgramFiles, RootProgramFilesX86:
		return true
	default:
		return false
	}
}

func (r Root) ValidForOS(os string) bool {
	switch r {
	case RootProcessPath, RootUserHome:
		return os == "darwin" || os == "linux" || os == "windows"
	case RootLoginShellPath:
		return os == "darwin" || os == "linux"
	case RootSystemApplications, RootUserApplications:
		return os == "darwin"
	case RootLocalAppData, RootAppData, RootProgramFiles, RootProgramFilesX86:
		return os == "windows"
	default:
		return false
	}
}

func portableAbsolute(value string) bool {
	return strings.HasPrefix(value, "/") || strings.HasPrefix(value, `\`) ||
		(len(value) >= 2 && ((value[0] >= 'A' && value[0] <= 'Z') || (value[0] >= 'a' && value[0] <= 'z')) && value[1] == ':')
}

func boundedID(value string, limit int) bool {
	value = strings.TrimSpace(value)
	if value == "" || len(value) > limit {
		return false
	}
	for _, r := range value {
		if !strings.ContainsRune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789._:-", r) {
			return false
		}
	}
	return true
}
