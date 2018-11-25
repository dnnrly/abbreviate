package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatcher_Match(t *testing.T) {
	m := Matcher{items: map[string]string{"abbreviation": "abbr"}}

	assert.Equal(t, "abbr", m.Match("abbreviation"))
	assert.Equal(t, "something", m.Match("something"))
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
