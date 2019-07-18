package domain

import (
	"fmt"
	"sort"
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
