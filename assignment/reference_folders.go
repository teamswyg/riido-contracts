package assignment

import (
	"encoding/json"
	"errors"
	"strings"
)

const MaxReferenceFolders = 20

// ReferenceFolders binds local paths to the runtime that can interpret them.
type ReferenceFolders struct {
	DeviceID  string   `json:"device_id"`
	RuntimeID string   `json:"runtime_id"`
	Paths     []string `json:"paths"`
}

func (r *ReferenceFolders) Clone() *ReferenceFolders {
	if r == nil {
		return nil
	}
	copy := *r
	copy.Paths = append([]string{}, r.Paths...)
	return &copy
}

// Normalize validates syntax only; local filesystem access belongs to daemon.
func (r *ReferenceFolders) Normalize() (*ReferenceFolders, error) {
	if r == nil {
		return nil, nil
	}
	if len(r.Paths) > MaxReferenceFolders {
		return nil, errors.New("reference_folders supports at most 20 paths")
	}
	if strings.TrimSpace(r.DeviceID) == "" || strings.TrimSpace(r.RuntimeID) == "" {
		return nil, errors.New("reference_folders requires device_id and runtime_id")
	}
	result := &ReferenceFolders{DeviceID: r.DeviceID, RuntimeID: r.RuntimeID, Paths: []string{}}
	seen := map[string]bool{}
	for _, path := range r.Paths {
		windowsDrive := len(path) >= 3 && ((path[0] >= 'A' && path[0] <= 'Z') || (path[0] >= 'a' && path[0] <= 'z')) && path[1] == ':' && (path[2] == '\\' || path[2] == '/')
		if len(path) > 4096 || strings.ContainsRune(path, 0) || !(strings.HasPrefix(path, "/") || strings.HasPrefix(path, `\\`) || windowsDrive) {
			return nil, errors.New("reference_folders paths must be absolute paths of at most 4096 bytes")
		}
		if !seen[path] {
			result.Paths = append(result.Paths, path)
			seen[path] = true
		}
	}
	return result, nil
}

func (r *ReferenceFolders) Matches(deviceID, runtimeID string) bool {
	return r != nil && r.DeviceID == deviceID && r.RuntimeID == runtimeID
}

func (r *ReferenceFolders) Instruction() string {
	if r == nil || len(r.Paths) == 0 {
		return ""
	}
	paths, _ := json.Marshal(r.Paths)
	return "[Reference folders]\nThe user selected these local folders as reference material (JSON paths): " + string(paths) +
		"\nAt the start of this task, inspect these folders and read material relevant to the request. Report any path you cannot access; never claim to have read material without actually reading it. These paths do not change your execution cwd. Folder selection alone does not request edits; follow the user's task when deciding whether to modify files."
}
