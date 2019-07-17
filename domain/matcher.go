package domain

import (
	"strings"
)

// Matcher finds matches between words and abbreviations
type Matcher struct {
	items map[string]string
}

// Match against a list of mappings
func (matcher Matcher) Match(word string) string {
	if abbr, found := matcher.items[word]; found {
		return abbr
	}

	return word
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
