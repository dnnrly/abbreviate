package domain

import (
	"reflect"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
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
		frmFront bool
		rmvStop  bool
		want     string
	}{
		{name: "Length longer than origin with '-'", original: "aaa-bbb-ccc", max: 99, want: "aaa-bbb-ccc"},
		{name: "Length is 0 with '-'", original: "aaa-bbb-ccc", max: 0, want: "a-b-c"},
		{name: "Partial abbreviation with '-'", original: "aaa-bbb-ccc", max: 10, want: "aaa-bbb-c"},
		{name: "Partial abbreviation with '-', start from the front", original: "aaa-bbb-ccc", max: 10, frmFront: true, want: "a-bbb-ccc"},
		{name: "Length longer than origin with camel case", original: "AaaBbbCcc", max: 99, want: "AaaBbbCcc"},
		{name: "Length is 0 with camel case", original: "AaaBbbCcc", max: 0, want: "ABC"},
		{name: "Length is 0 with camel case, matching case", original: "aaaBbbCcc", max: 0, want: "aBC"},
		{name: "Partial abbreviation with camel case", original: "AaaBbbCcc", max: 8, want: "AaaBbbC"},
		{name: "Partial abbreviation with camel case, start from the front", original: "AaaBbbCcc", max: 8, frmFront: true, want: "ABbbCcc"},
		{name: "Doesn't match wrong casing", original: "AaaBBbCcc", max: 0, want: "ABBbC"},
		{name: "Mixed camel case and non word separators", original: "AaaBbb-ccc", max: 0, want: "AB-c"},
		{name: "Mixed camel case and non word separators with same borders", original: "Aaa-Bbb-Ccc", max: 0, want: "A-B-C"},
		{name: "Real example, full short", original: "strategy-limited", max: 0, want: "stg-ltd"},
		{name: "Real example, shorter than total", original: "strategy-limited", max: 13, want: "strategy-ltd"},
		{name: "Real example, shorter than total, start from the front", original: "strategy-limited", max: 13, frmFront: true, want: "stg-limited"},
		{name: "Real example, max same as shorted", original: "strategy-limited", max: 12, want: "strategy-ltd"},
		{name: "Real example, max same as shorted, start from the front", original: "strategy-limited", max: 12, frmFront: true, want: "stg-limited"},
		{name: "Real example, max on separator", original: "strategy-limited", max: 9, want: "stg-ltd"},
		{name: "Real example, max shorter than first word", original: "strategy-limited", max: 6, want: "stg-ltd"},
		{name: "Real example, no short", original: "strategy-limited", max: 99, want: "strategy-limited"},
		{name: "Real example, with numbers #1", original: "strategy-limited99", max: 15, want: "strategy-ltd99"},
		{name: "Real example, with numbers #2", original: "strategy-limited-99", max: 15, want: "strategy-ltd-99"},
		{name: "Real example, with numbers, start from the front #1", original: "strategy-limited99", max: 15, frmFront: true, want: "stg-limited99"},
		{name: "Real example, with numbers, start from the front #2", original: "strategy-limited-99", max: 15, frmFront: true, want: "stg-limited-99"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AsOriginal(matcher, tt.original, tt.max, tt.frmFront, tt.rmvStop); got != tt.want {
				t.Errorf("AsOriginal('%s', %d) = '%v', want '%v'", tt.original, tt.max, got, tt.want)
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

func TestSequences_AddFront(t *testing.T) {
	seqs := Sequences{}

	seqs.AddFront("a")
	seqs.AddFront("b")
	seqs.AddFront("cd")

	assert.Equal(t, 3, len(seqs))
	assert.Equal(t, "cd", seqs[0])
	assert.Equal(t, "b", seqs[1])
	assert.Equal(t, "a", seqs[2])
	assert.Equal(t, "cdba", seqs.String())
	assert.Equal(t, 4, seqs.Len())
}

func TestSequences_AddBack(t *testing.T) {
	seqs := Sequences{}

	seqs.AddBack("a")
	seqs.AddBack("b")
	seqs.AddBack("cd")

	assert.Equal(t, 3, len(seqs))
	assert.Equal(t, "a", seqs[0])
	assert.Equal(t, "b", seqs[1])
	assert.Equal(t, "cd", seqs[2])
	assert.Equal(t, "abcd", seqs.String())
}

func TestNewSequences(t *testing.T) {
	tests := []struct {
		name string
		orig string
		want Sequences
	}{
		{name: "1", orig: "abc", want: Sequences{"abc"}},
		{name: "2", orig: "a-b-c", want: Sequences{"a", "-", "b", "-", "c"}},
		{name: "3", orig: "ABC", want: Sequences{"A", "B", "C"}},
		{name: "4", orig: "a-b--c", want: Sequences{"a", "-", "b", "--", "c"}},
		{name: "5", orig: "aa-bb", want: Sequences{"aa", "-", "bb"}},
		{name: "6", orig: "aaBbCc", want: Sequences{"aa", "Bb", "Cc"}},
		{name: "7", orig: "aa-Bb-cc", want: Sequences{"aa", "-", "Bb", "-", "cc"}},
		{name: "8", orig: "AaaBBbCcc", want: Sequences{"Aaa", "B", "Bb", "Ccc"}},
		{name: "9", orig: "AaaBBb888Ccc", want: Sequences{"Aaa", "B", "Bb", "888", "Ccc"}},
		{name: "10", orig: "AaaBBb-8Ccc", want: Sequences{"Aaa", "B", "Bb", "-", "8", "Ccc"}},
		{name: "11", orig: "", want: Sequences{}},
		{name: "12", orig: "a", want: Sequences{"a"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSequences(tt.orig); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSequences() = %v (%d), want %v (%d)",
					[]string(got), len(got),
					[]string(tt.want), len(tt.want),
				)
			}
		})
	}
}

func Test_isTitleCase(t *testing.T) {
	assert.Equal(t, true, isTitleCase("Abc"))
	assert.Equal(t, true, isTitleCase("A"))
	assert.Equal(t, true, isTitleCase("ABC"))
	assert.Equal(t, false, isTitleCase("abc"))
	assert.Equal(t, false, isTitleCase("aBC"))
	assert.Equal(t, false, isTitleCase("a"))
}

func Test_first(t *testing.T) {
	assert.Equal(t, rune(0), first(""))
	assert.Equal(t, rune('a'), first("a"))
	assert.Equal(t, rune('c'), first("cba"))
	assert.Equal(t, rune('B'), first("Bac"))
	assert.Equal(t, false, unicode.IsTitle(first("")))
}

func TestAsSeparated(t *testing.T) {
	matcher := NewMatcherFromString(`a=aaa
b=bbb
c=ccc
d=ddd
stg=strategy
ltd=limited`)
	tests := []struct {
		name      string
		original  string
		separator string
		max       int
		frmFront  bool
		rmvStop   bool
		want      string
	}{
		{name: "Length is 0 with '-'", original: "aaa-bbb-ccc", separator: "_", max: 0, want: "a_b_c"},
		{name: "Partial abbreviation with '_'", original: "aaa-bbb-ccc", separator: "_", max: 10, want: "aaa_bbb_c"},
		{name: "Length longer than origin with camel case", original: "AaaBbbCcc", separator: "_", max: 99, want: "aaa_bbb_ccc"},
		{name: "Length is 0 with camel case", original: "AaaBbbCcc", separator: "_", max: 0, want: "a_b_c"},
		{name: "Length is 0 with camel case, matching case", original: "aaaBbbCcc", separator: "_", max: 0, want: "a_b_c"},
		{name: "Partial abbreviation with camel case", original: "AaaBbbCcc", separator: "_", max: 8, want: "aaa_b_c"},
		{name: "Partial abbreviation with camel case, start from the front", original: "AaaBbbCcc", separator: "_", max: 8, frmFront: true, want: "a_b_ccc"},
		{name: "Doesn't match wrong casing", original: "AaaBBbCcc", separator: "_", max: 0, want: "a_b_bb_c"},
		{name: "Mixed camel case and non word separators", original: "AaaBbb-ccc", separator: "_", max: 0, want: "a_b_c"},
		{name: "Mixed camel case and non word separators with same borders", separator: "_", original: "Aaa-Bbb-Ccc", max: 0, want: "a_b_c"},
		{name: "Snake case, full short", original: "strategy_limited", separator: "-", max: 0, want: "stg-ltd"},
		{name: "Real example, full short", original: "strategy-limited", separator: "_", max: 0, want: "stg_ltd"},
		{name: "Real example, shorter than total", original: "strategy-limited", separator: "_", max: 13, want: "strategy_ltd"},
		{name: "Real example, shorter than total, start from the front", original: "strategy-limited", separator: "_", max: 13, frmFront: true, want: "stg_limited"},
		{name: "Real example, max same as shorted", original: "strategy-limited", separator: "_", max: 12, want: "strategy_ltd"},
		{name: "Real example, max same as shorted, start from the front", original: "strategy-limited", separator: "_", max: 12, frmFront: true, want: "stg_limited"},
		{name: "Real example, max on separator", original: "strategy-limited", separator: "_", max: 9, want: "stg_ltd"},
		{name: "Real example, max shorter than first word", original: "strategy-limited", separator: "_", max: 6, want: "stg_ltd"},
		{name: "Real example, no short", original: "strategy-limited", separator: "_", max: 99, want: "strategy_limited"},
		{name: "Real example, with numbers #1", original: "strategy-limited99", separator: "_", max: 15, want: "strategy_ltd_99"},
		{name: "Real example, with numbers #2", original: "strategy-limited-99", separator: "_", max: 15, want: "strategy_ltd_99"},
		{name: "Real example, with numbers, start from the front #1", original: "strategy-limited99", separator: "_", max: 15, frmFront: true, want: "stg_limited_99"},
		{name: "Real example, with numbers, start from the front #2", original: "strategy-limited-99", separator: "_", max: 15, frmFront: true, want: "stg_limited_99"},
		{name: "Multiple separators", original: "strategy---limited--99", separator: "_", max: 15, want: "strategy_ltd_99"},
		{name: "Multiple separators, start from the front", original: "strategy---limited--99", separator: "_", max: 15, frmFront: true, want: "stg_limited_99"},
		{name: "Other separator", original: "strategy-limited-99", separator: "+", max: 15, want: "strategy+ltd+99"},
		{name: "Other separator, start from the front", original: "strategy-limited-99", separator: "+", max: 15, frmFront: true, want: "stg+limited+99"},
		{name: "Empty string", original: "", separator: "+", max: 15, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AsSeparated(matcher, tt.original, tt.separator, tt.max, tt.frmFront, tt.rmvStop); got != tt.want {
				t.Errorf("AsSeparated() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAsPascal(t *testing.T) {
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
		frmFront bool
		rmvStop  bool
		want     string
	}{
		{name: "Length longer than origin with '-'", original: "aaa-bbb-ccc", max: 99, want: "AaaBbbCcc"},
		{name: "Length is 0 with '-'", original: "aaa-bbb-ccc", max: 0, want: "ABC"},
		{name: "Partial abbreviation with '-'", original: "aaa-bbb-ccc", max: 8, want: "AaaBbbC"},
		{name: "Partial abbreviation with '-', start from the front", original: "aaa-bbb-ccc", max: 8, frmFront: true, want: "ABbbCcc"},
		{name: "Length longer than origin with camel case", original: "AaaBbbCcc", max: 99, want: "AaaBbbCcc"},
		{name: "Length is 0 with camel case", original: "AaaBbbCcc", max: 0, want: "ABC"},
		{name: "Length is 0 with camel case, matching case", original: "aaaBbbCcc", max: 0, want: "ABC"},
		{name: "Partial abbreviation with camel case", original: "AaaBbbCcc", max: 8, want: "AaaBbbC"},
		{name: "Partial abbreviation with camel case, start from the front", original: "AaaBbbCcc", max: 8, frmFront: true, want: "ABbbCcc"},
		{name: "Doesn't match wrong casing", original: "AaaBBbCcc", max: 0, want: "ABBbC"},
		{name: "Mixed camel case and non word separators", original: "AaaBbb-ccc", max: 0, want: "ABC"},
		{name: "Mixed camel case and non word separators with same borders", original: "Aaa-Bbb-Ccc", max: 0, want: "ABC"},
		{name: "Real example, full short", original: "strategy-limited", max: 0, want: "StgLtd"},
		{name: "Real example, shorter than total", original: "strategy-limited", max: 13, want: "StrategyLtd"},
		{name: "Real example, shorter than total, start from the front", original: "strategy-limited", max: 13, frmFront: true, want: "StgLimited"},
		{name: "Real example, max same as shorted", original: "strategy-limited", max: 12, want: "StrategyLtd"},
		{name: "Real example, max same as shorted", original: "strategy-limited", max: 12, frmFront: true, want: "StgLimited"},
		{name: "Real example, max on separator", original: "strategy-limited", max: 9, want: "StgLtd"},
		{name: "Real example, max shorter than first word", original: "strategy-limited", max: 6, want: "StgLtd"},
		{name: "Real example, no short", original: "strategy-limited", max: 99, want: "StrategyLimited"},
		{name: "Real example, with numbers #1", original: "strategy-limited99", max: 15, want: "StrategyLtd99"},
		{name: "Real example, with numbers #2", original: "strategy-limited-99", max: 15, want: "StrategyLtd99"},
		{name: "Real example, with numbers #1, start from the front", original: "strategy-limited99", max: 15, frmFront: true, want: "StgLimited99"},
		{name: "Real example, with numbers #2, start from the front", original: "strategy-limited-99", max: 15, frmFront: true, want: "StgLimited99"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AsPascal(matcher, tt.original, tt.max, tt.frmFront, tt.rmvStop); got != tt.want {
				t.Errorf("AsPascal('%s', %d) = '%v', want '%v'", tt.original, tt.max, got, tt.want)
			}
		})
	}
}
