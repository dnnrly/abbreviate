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

func TestMatcher_Prefixes(t *testing.T) {

	m := Matcher{mainWords: map[string]string{"abbreviation": "abbr"},
		prefixes: map[string]string{"pre": "pr", "anti": "ant"}}

	assert.Equal(t, "prabbr", m.Match("preabbreviation"))
	assert.Equal(t, "antprabbr", m.Match("antipreabbreviation"))
}

func TestMatcher_Suffixes(t *testing.T) {

	m := Matcher{mainWords: map[string]string{"censor": "cnsr"},
		suffixes: map[string]string{"ship": "shp"}}

	assert.Equal(t, "cnsrshp", m.Match("censorship"))
}

func TestMatcher_PrefixesSuffixes(t *testing.T) {

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
