package capability

import (
	"encoding/json"
	"fmt"
	"sort"
)

func marshalSortedMap(m map[string]any) ([]byte, error) {
	if m == nil {
		return []byte("[]"), nil
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	out := make([]sortedFlag, 0, len(keys))
	for _, k := range keys {
		out = append(out, sortedFlag{K: k, V: m[k]})
	}
	data, err := json.Marshal(out)
	if err != nil {
		return nil, fmt.Errorf("canonicalize surface flags: %w", err)
	}
	return data, nil
}

type sortedFlag struct {
	K string `json:"k"`
	V any    `json:"v"`
}
