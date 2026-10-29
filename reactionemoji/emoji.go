// Package reactionemoji preserves the v1 unified-code reaction contract across
// Riido storage and Slack, including glyph values written by earlier clients.
package reactionemoji

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode/utf8"
)

//go:embed catalog.tsv
var catalog string

//go:embed skin-variants.txt
var skinVariants string

type catalogEntry struct{ unified, slack string }

var byCode, byName = loadCatalog()
var skinCodes = loadSkinCodes()

func loadSkinCodes() map[string]struct{} {
	codes := make(map[string]struct{})
	for _, code := range strings.Fields(skinVariants) {
		codes[identity(code)] = struct{}{}
	}
	return codes
}

func loadCatalog() (map[string]catalogEntry, map[string]string) {
	codes := make(map[string]catalogEntry)
	names := make(map[string]string)
	for _, line := range strings.Split(strings.TrimSpace(catalog), "\n") {
		fields := strings.Split(line, "\t")
		if len(fields) != 4 {
			panic("invalid embedded v1 reaction emoji catalog")
		}
		entry := catalogEntry{unified: fields[0], slack: fields[2]}
		codes[identity(entry.unified)] = entry
		for _, name := range strings.Split(fields[3], ",") {
			names[name] = entry.unified
		}
	}
	// Preserve names accepted by the previous inbound adapter.
	names["red_heart"] = names["heart"]
	return codes, names
}

func identity(code string) string {
	parts := strings.Split(code, "-")
	parts = slices.DeleteFunc(parts, func(part string) bool { return part == "fe0f" })
	return strings.Join(parts, "-")
}

// Normalize returns lowercase unified codepoints. Known presentation variants
// use the fixed v1 catalog's qualified value. Unrecognized text is left intact.
func Normalize(value string) string {
	code, ok := codepoints(value)
	if !ok {
		return value
	}
	if entry, found := byCode[identity(code)]; found {
		return entry.unified
	}
	return code
}

func codepoints(value string) (string, bool) {
	value = strings.TrimSpace(value)
	if value == "" {
		return "", false
	}
	parts := strings.Split(value, "-")
	hexadecimal := true
	for _, part := range parts {
		if part == "" {
			hexadecimal = false
			break
		}
		for _, r := range part {
			if !strings.ContainsRune("0123456789abcdefABCDEF", r) {
				hexadecimal = false
				break
			}
		}
		if !hexadecimal {
			break
		}
	}
	if hexadecimal {
		for index, part := range parts {
			r, err := strconv.ParseUint(part, 16, 32)
			if err != nil || !utf8.ValidRune(rune(r)) || r < 0x20 {
				return "", false
			}
			parts[index] = fmt.Sprintf("%04x", r)
		}
	} else {
		if !utf8.ValidString(value) {
			return "", false
		}
		parts = nil
		for _, r := range value {
			parts = append(parts, fmt.Sprintf("%04x", r))
		}
	}
	code := strings.Join(parts, "-")
	if _, known := byCode[identity(code)]; known {
		return code, true
	}
	// Nested v1 skin variants are recognized without inventing Slack names or
	// changing their existing presentation form.
	_, known := skinCodes[identity(code)]
	return code, known
}

// SlackName uses every top-level emoji supported by the fixed v1 catalog.
func SlackName(value string) string {
	return byCode[identity(Normalize(value))].slack
}

// FromSlackName returns the same lowercase unified value used by Riido clients.
// Unknown/custom Slack emoji remain unsupported, as in v1.
func FromSlackName(name string) string { return byName[strings.ToLower(name)] }

// Aliases supplies legacy glyph/unqualified and code forms for storage lookups.
// These aliases are exact matches so unrecognized text retains its case.
func Aliases(value string) []string {
	canonical := Normalize(value)
	result := []string{value}
	for _, code := range codeAliases(canonical) {
		result = append(result, code)
		if glyph, ok := glyph(code); ok && Normalize(glyph) == canonical {
			result = append(result, glyph)
		}
	}
	slices.Sort(result)
	return slices.Compact(result)
}

// CodeAliases supplies only recognized hexadecimal representations for SQL
// case folding; unknown text and glyphs must use exact Aliases matches.
func CodeAliases(value string) []string {
	if _, ok := codepoints(value); !ok {
		return nil
	}
	return codeAliases(Normalize(value))
}

func codeAliases(canonical string) []string {
	codes := []string{canonical}
	if _, known := byCode[identity(canonical)]; known {
		codes = []string{""}
		parts := strings.Split(canonical, "-")
		for index := 0; index < len(parts); index++ {
			part := parts[index]
			_, standalone := byCode[part]
			// Only catalog emoji bases can gain a presentation selector. Do not
			// insert one after joiners, flag/tag components, or skin modifiers.
			presentation := standalone && !(part >= "1f3fb" && part <= "1f3ff")
			if index+1 < len(parts) && parts[index+1] == "fe0f" {
				presentation = true
				index++
			}
			next := make([]string, 0, len(codes)*2)
			for _, prefix := range codes {
				joined := part
				if prefix != "" {
					joined = prefix + "-" + part
				}
				next = append(next, joined)
				if presentation {
					next = append(next, joined+"-fe0f")
				}
			}
			codes = next
		}
	}
	return codes
}

func glyph(code string) (string, bool) {
	var result strings.Builder
	for _, part := range strings.Split(code, "-") {
		r, err := strconv.ParseUint(part, 16, 32)
		if err != nil || !utf8.ValidRune(rune(r)) || r < 0x20 {
			return "", false
		}
		result.WriteRune(rune(r))
	}
	return result.String(), true
}
