package domain

import (
	"testing"
)

func TestShortenFromBack(t *testing.T) {
	matcher := NewMatcherFromString(`a=aaa
b=bbb
c=ccc
d=ddd
stg=strategy
ltd=limited`)
	tests := []struct {
		name     string
		original string
		max      int
		want     string
	}{
		{name: "Length longer than origin with '-'", original: "aaa-bbb-ccc", max: 99, want: "aaa-bbb-ccc"},
		{name: "Length is 0 with '-'", original: "aaa-bbb-ccc", max: 0, want: "a-b-c"},
		{name: "Partial abbreviation with '-'", original: "aaa-bbb-ccc", max: 10, want: "aaa-bbb-c"},
		{name: "Length longer than origin with camel case", original: "AaaBbbCcc", max: 99, want: "AaaBbbCcc"},
		{name: "Length is 0 with camel case", original: "AaaBbbCcc", max: 0, want: "ABC"},
		{name: "Length is 0 with camel case, matching case", original: "aaaBbbCcc", max: 0, want: "aBC"},
		{name: "Partial abbreviation with camel case", original: "AaaBbbCcc", max: 8, want: "AaaBbbC"},
		{name: "Doesn't match wrong casing", original: "AaaBBbCcc", max: 0, want: "ABBbC"},
		{name: "Mixed camel case and non word seperators", original: "AaaBbb-ccc", max: 0, want: "AB-c"},
		{name: "Mixed camel case and non word seperators with same borders", original: "Aaa-Bbb-Ccc", max: 0, want: "A-B-C"},
		{name: "Real example, full short", original: "strategy-limited", max: 0, want: "stg-ltd"},
		{name: "Real example, shorter than total", original: "strategy-limited", max: 13, want: "strategy-ltd"},
		{name: "Real example, max same as shorted", original: "strategy-limited", max: 12, want: "strategy-ltd"},
		{name: "Real example, max on seperator", original: "strategy-limited", max: 9, want: "strategy-ltd"},
		{name: "Real example, max shorter than first word", original: "strategy-limited", max: 6, want: "stg-ltd"},
		{name: "Real example, no short", original: "strategy-limited", max: 99, want: "strategy-limited"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShortenFromBack(matcher, tt.original, tt.max); got != tt.want {
				t.Errorf("ShortenFromBack('%s', %d) = '%v', want '%v'", tt.original, tt.max, got, tt.want)
			}
		})
	}
}

func Test_lastChar(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		remains string
		last    rune
	}{
		{name: "1", str: "string", remains: "strin", last: 'g'},
		{name: "2", str: "s", remains: "", last: 's'},
		{name: "3", str: "", remains: "", last: rune(0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := lastChar(tt.str)
			if got != tt.remains {
				t.Errorf("lastChar() got = '%v', want '%v'", got, tt.remains)
			}
			if got1 != tt.last {
				t.Errorf("lastChar() got1 = %v, want %v", got1, tt.last)
			}
		})
	}
}
