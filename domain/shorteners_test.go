package domain

import "testing"

func TestShortenFromBack(t *testing.T) {
	matcher := NewMatcherFromString(`a=aaa
b=bbb
c=ccc
d=ddd`)
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShortenFromBack(matcher, tt.original, tt.max); got != tt.want {
				t.Errorf("ShortenFromBack('%s', %d) = '%v', want '%v'", tt.original, tt.max, got, tt.want)
			}
		})
	}
}
