package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatcher_Match(t *testing.T) {
	m := Matcher{mainWords: map[string]string{"abbreviation": "abbr"}}

	assert.Equal(t, "abbr", m.Match("abbreviation"))
	assert.Equal(t, "something", m.Match("something"))
}

func TestMatcher_Match_prefixess(t *testing.T) {
	m := Matcher{mainWords: map[string]string{"abbreviation": "abbr"},
		prefixes: map[string]string{"pre": "pr", "anti": "ant"}}

	assert.Equal(t, "prabbr", m.Match("preabbreviation"))
	assert.Equal(t, "antprabbr", m.Match("antipreabbreviation"))
}

func TestMatcher_Match_suffixes(t *testing.T) {
	m := Matcher{mainWords: map[string]string{"censor": "cnsr"},
		suffixes: map[string]string{"ship": "shp"}}

	assert.Equal(t, "cnsrshp", m.Match("censorship"))
}

func TestMatcher_Match_multiple_prefixes_suffixes(t *testing.T) {
	m := Matcher{mainWords: map[string]string{"establish": "estblsh"},
		prefixes: map[string]string{
			"anti": "ant",
			"dis":  "ds"},
		suffixes: map[string]string{
			"ment":  "mnt",
			"arian": "arn",
			"ism":   "sm"}}

	assert.Equal(t, "antdsestblshmntarnsm", m.Match("antidisestablishmentarianism"))
}

func TestMatcher_Match_partial_match_prefixes_suffixes(t *testing.T) {
	m := Matcher{
		mainWords: map[string]string{"establish": "estblsh"},
		prefixes: map[string]string{
			"anti": "ant",
		},
		suffixes: map[string]string{
			"ment":  "mnt",
			"arian": "arn",
		},
	}

	// Prefix and suffix checking stops recursing as soon as there is no match.
	// Here, 'ment' and 'arian' are never abbreviated because 'ism' is not matched.
	assert.Equal(t, "antdisestablishmentarianism", m.Match("antidisestablishmentarianism"))
}

func TestMatcher_no_prefix_suffix_returns_original_word(t *testing.T) {
	m := Matcher{mainWords: map[string]string{"establish": "estblsh"}}
	assert.Equal(t, "establishment", m.Match("establishment"))
}

func TestMatcher_NewMatcherFromString(t *testing.T) {
	m := NewMatcherFromString(`abbr=abbreviation
1=a
2=b

3
=c
4=d=dd

`)

	assert.Equal(t, "abbr", m.Match("abbreviation"))
	assert.Equal(t, "2", m.Match("b"))
	assert.Equal(t, "3", m.Match("3"))
	assert.Equal(t, "c", m.Match("c"))
	assert.Equal(t, "d", m.Match("d"))
	assert.Equal(t, "dd", m.Match("dd"))
}

func TestMatcher_All(t *testing.T) {
	m := NewMatcherFromString(`abbr=abbreviation
1=a
2=b

3
=c
4=d=dd

`)
	expected := []string{
		"1=a",
		"2=b",
		"abbr=abbreviation",
	}
	assert.Equal(t, expected, m.All())
}
