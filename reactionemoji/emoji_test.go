package reactionemoji

import (
	"slices"
	"strings"
	"testing"
)

func TestEntireV1CatalogRoundTrip(t *testing.T) {
	lines := strings.Split(strings.TrimSpace(catalog), "\n")
	if len(lines) != 1903 {
		t.Fatalf("v1 catalog entries=%d, want 1903", len(lines))
	}
	for _, line := range lines {
		fields := strings.Split(line, "\t")
		unified, unqualified, name := fields[0], fields[1], fields[2]
		for _, code := range []string{unified, unqualified} {
			if code == "" {
				continue
			}
			unicode, ok := glyph(code)
			if !ok {
				t.Fatalf("invalid catalog code %q", code)
			}
			for _, input := range []string{code, strings.ToUpper(code), unicode} {
				if Normalize(input) != unified || SlackName(input) != name {
					t.Errorf("catalog round trip %q: normalized=%q slack=%q; want %q/%q", input, Normalize(input), SlackName(input), unified, name)
				}
			}
		}
		for _, alias := range strings.Split(fields[3], ",") {
			if got := FromSlackName(alias); got != unified {
				t.Errorf("Slack alias %q=%q want %q", alias, got, unified)
			}
		}
	}
}

func TestV1SkinVariantsPreservePresentationWithoutSlackNames(t *testing.T) {
	codes := strings.Fields(skinVariants)
	if len(codes) != 1875 {
		t.Fatalf("v1 skin variants=%d, want 1875", len(codes))
	}
	for _, code := range codes {
		for _, presentation := range []string{code, identity(code)} {
			unicode, ok := glyph(presentation)
			if !ok {
				t.Fatalf("invalid skin variant %q", presentation)
			}
			for _, input := range []string{presentation, strings.ToUpper(presentation), unicode} {
				if got := Normalize(input); got != presentation {
					t.Errorf("Normalize(%q)=%q, want preserved presentation %q", input, got, presentation)
				}
				if got := CodeAliases(input); !slices.Equal(got, []string{presentation}) {
					t.Errorf("CodeAliases(%q)=%q, want only %q", input, got, presentation)
				}
				if got := SlackName(input); got != "" {
					t.Errorf("SlackName(%q)=%q, skin variants have no v1 Slack name", input, got)
				}
			}
		}
	}
}

func TestV1ReactionEmojiContract(t *testing.T) {
	cases := []struct{ input, normalized, slack string }{
		{"2705", "2705", "white_check_mark"},
		{"✅", "2705", "white_check_mark"},
		{"1F440", "1f440", "eyes"},
		{"👀", "1f440", "eyes"},
		{"1f60d", "1f60d", "heart_eyes"},
		{"😍", "1f60d", "heart_eyes"},
		{"❤", "2764-fe0f", "heart"},
		{"2764", "2764-fe0f", "heart"},
		{"❤️", "2764-fe0f", "heart"},
		{"👨‍💻", "1f468-200d-1f4bb", "male-technologist"},
		{"🇰🇷", "1f1f0-1f1f7", "kr"},
		{"1️⃣", "0031-fe0f-20e3", "one"},
	}
	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			if got := Normalize(tc.input); got != tc.normalized {
				t.Fatalf("Normalize(%q)=%q; want %q", tc.input, got, tc.normalized)
			}
			if got := SlackName(tc.input); got != tc.slack {
				t.Fatalf("SlackName(%q)=%q; want %q", tc.input, got, tc.slack)
			}
			if got := FromSlackName(tc.slack); got != tc.normalized {
				t.Fatalf("FromSlackName(%q)=%q; want %q", tc.slack, got, tc.normalized)
			}
		})
	}
}

func TestExistingReactionAliases(t *testing.T) {
	for _, value := range []string{"2764", "2764-fe0f", "❤", "❤️"} {
		aliases := Aliases(value)
		for _, want := range []string{"2764", "2764-fe0f", "❤", "❤️"} {
			if !slices.Contains(aliases, want) {
				t.Errorf("Aliases(%q)=%q missing %q", value, aliases, want)
			}
		}
	}
	if got := FromSlackName("thumbsup"); got != "1f44d" {
		t.Errorf("thumbsup alias = %q", got)
	}
	if got := FromSlackName("red_heart"); got != "2764-fe0f" {
		t.Errorf("red_heart alias = %q", got)
	}
}

func TestUnknownReactionAliasesRemainCaseSensitive(t *testing.T) {
	for _, value := range []string{"Foo", "foo", "A", "a", "D800", "d800", "FACE", "face", "CAFE", "cafe", "BEEF", "beef", "23AB", "23ab", "1F000", "1f000", "0041", "0061", "0030"} {
		if got := Aliases(value); !slices.Equal(got, []string{value}) {
			t.Errorf("Aliases(%q)=%q, want only the original value", value, got)
		}
		if got := CodeAliases(value); len(got) != 0 {
			t.Errorf("CodeAliases(%q)=%q, unknown text must not match case-insensitively", value, got)
		}
	}
}

func TestCodeAliasesExcludeGlyphs(t *testing.T) {
	for _, value := range []string{"1F44D-FE0F", "👍️"} {
		if got := CodeAliases(value); !slices.Equal(got, []string{"1f44d", "1f44d-fe0f"}) {
			t.Errorf("CodeAliases(%q)=%q, want only lowercase code variants", value, got)
		}
	}
	if got := CodeAliases("0061"); len(got) != 0 {
		t.Errorf("CodeAliases(0061)=%q, text glyph must not be case folded", got)
	}
}

func TestGlyphAliasesPreserveReadIdentity(t *testing.T) {
	for _, value := range []string{"0061", "0041", "0066-006f-006f", "0030", "0066-0061-0063-0065"} {
		canonical := Normalize(value)
		for _, alias := range Aliases(value) {
			if got := Normalize(alias); got != canonical {
				t.Errorf("Aliases(%q) contains %q with a different read identity %q", value, alias, got)
			}
		}
	}
}

func TestEmojiDefaultPresentationAliases(t *testing.T) {
	cases := []struct {
		canonical string
		aliases   []string
	}{
		{"1f44d", []string{"1f44d", "1f44d-fe0f", "👍", "👍️"}},
		{"1f468-200d-1f4bb", []string{"1f468-fe0f-200d-1f4bb", "1f468-200d-1f4bb-fe0f", "1f468-fe0f-200d-1f4bb-fe0f", "👨️‍💻️"}},
		{"1f3cb-fe0f-200d-2640-fe0f", []string{"1f3cb-200d-2640", "1f3cb-fe0f-200d-2640", "1f3cb-200d-2640-fe0f", "🏋‍♀"}},
	}
	for _, tc := range cases {
		t.Run(tc.canonical, func(t *testing.T) {
			for _, input := range append([]string{tc.canonical}, tc.aliases...) {
				if got := Normalize(input); got != tc.canonical {
					t.Fatalf("Normalize(%q)=%q, want %q", input, got, tc.canonical)
				}
				aliases := Aliases(input)
				for _, want := range tc.aliases {
					if !slices.Contains(aliases, want) {
						t.Errorf("Aliases(%q)=%q missing %q", input, aliases, want)
					}
				}
			}
		})
	}
}

func TestPresentationAliasesPreserveSequenceIdentity(t *testing.T) {
	for _, line := range strings.Split(strings.TrimSpace(catalog), "\n") {
		canonical := strings.Split(line, "\t")[0]
		for _, alias := range Aliases(canonical) {
			if got := Normalize(alias); got != canonical {
				t.Errorf("alias %q changed identity from %q to %q", alias, canonical, got)
			}
		}
	}
	for _, tc := range []struct {
		input string
		want  []string
	}{
		{"1f44d-1f3fb", []string{"1f44d-1f3fb", "👍🏻"}},
		{"1f3fb", []string{"1f3fb", "🏻"}},
		{"1f1f0-1f1f7", []string{"1f1f0-1f1f7", "🇰🇷"}},
		{"custom_emoji", []string{"custom_emoji"}},
	} {
		if got := Aliases(tc.input); !slices.Equal(got, tc.want) {
			t.Errorf("Aliases(%q)=%q, want %q", tc.input, got, tc.want)
		}
	}
	for _, tc := range []struct{ canonical, excluded string }{
		{"1f468-200d-1f4bb", "1f468-200d-fe0f-1f4bb"},
		{"1f3f4-e0067-e0062-e0065-e006e-e0067-e007f", "1f3f4-e0067-fe0f-e0062-e0065-e006e-e0067-e007f"},
		{"1f44d", "1f44d-fe0f-fe0f"},
	} {
		if slices.Contains(Aliases(tc.canonical), tc.excluded) {
			t.Errorf("Aliases(%q) contains invalid presentation %q", tc.canonical, tc.excluded)
		}
	}
}

func TestUnknownValuesPreserveStorageWithoutInventingSlackNames(t *testing.T) {
	for _, value := range []string{"custom_emoji", "not an emoji", "", "110000", "d800", "FACE", "face", "CAFE", "cafe", "BEEF", "beef", "23AB", "23ab", "1F000", "1f000"} {
		if got := Normalize(value); got != value {
			t.Errorf("Normalize(%q)=%q", value, got)
		}
		if got := SlackName(value); got != "" {
			t.Errorf("SlackName(%q)=%q", value, got)
		}
	}
	if got := Normalize("👍🏻"); got != "1f44d-1f3fb" {
		t.Errorf("skin sequence normalized to %q", got)
	}
	if got := Normalize("1F44D-1F3FB"); got != "1f44d-1f3fb" {
		t.Errorf("code sequence normalized to %q", got)
	}
	if got := FromSlackName("custom_emoji"); got != "" {
		t.Errorf("custom Slack emoji = %q", got)
	}
}
