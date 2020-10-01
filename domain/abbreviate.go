package domain

import (
	"fmt"
	"sort"
	"strings"
)

// Abbreviator allows you to use different implementations for abbreviating words
type Abbreviator interface {
	// Abbreviate attempts to return an abbreviation for the word given. It
	// does not guarantee to return a different, shorter string.
	Abbreviate(word string) string
}

// NoAbbreviator does not abbreviate anything
type NoAbbreviator struct{}

// Abbreviate just returns the unabbreviated word
func (a *NoAbbreviator) Abbreviate(word string) string {
	return word
}

// Matcher finds matches between words and abbreviations
type Matcher struct {
	items map[string]string
}

// NewMatcher creates a new matcher of abbreviation mappings
func NewMatcher(items map[string]string) *Matcher {
	return &Matcher{
		items: items,
	}
}

// Abbreviate returns an abbreviation looked up from pre-defined mappings
func (matcher *Matcher) Abbreviate(word string) string {
	return (*matcher).Match(word)
}

// Match against a list of mappings
func (matcher Matcher) Match(word string) string {
	if abbr, found := matcher.items[word]; found {
		return abbr
	}

	return word
}

// All of the abbreviations in this set in order of the linked word
func (matcher Matcher) All() []string {
	all := []string{}

	for k, v := range matcher.items {
		all = append(all, fmt.Sprintf("%s=%s", v, k))
	}

	sort.Strings(all)

	return all
}

// NewMatcherFromString creates a new matcher from newline
// seperated data
func NewMatcherFromString(data string) *Matcher {
	items := map[string]string{}

	lines := strings.Split(data, "\n")
	for _, line := range lines {
		parts := strings.Split(line, "=")
		if len(parts) == 2 {
			if parts[0] != "" {
				items[parts[1]] = parts[0]
			}
		}
	}

	return &Matcher{items: items}
}
