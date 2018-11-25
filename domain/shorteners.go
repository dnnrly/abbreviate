package domain

import (
	"strings"
	"unicode"
)

type parseState int

const (
	starting parseState = iota
	inWord
	betweenWords
)

// Shortener represents an algorithm that can be used to shorten a string
// by substituting words for abbreviations
type Shortener func(matcher Matcher, original string, max int) string

// ShortenFromBack discovers words using camel case and non letter characters,
// starting from the back until the string has less than 'max' characters
// or it can't shorten any more.
func ShortenFromBack(matcher *Matcher, original string, max int) string {
	if len(original) < max {
		return original
	}

	shortened := ""
	word := ""
	for pos := len(original) - 1; pos >= 0; pos-- {
		ch := []rune(original)[pos]
		if unicode.IsLetter(ch) {
			word = string(ch) + word
			if unicode.IsUpper(ch) {
				word = strings.ToLower(word)
				abbr := matcher.Match(word)
				abbr = strings.Title(abbr)
				shortened = abbr + shortened
				word = ""
			} else if pos == 0 {
				abbr := matcher.Match(word)
				shortened = abbr + shortened
			}
		} else {
			if word != "" {
				abbr := matcher.Match(word)
				shortened = abbr + shortened
				word = ""
			}
			shortened = string(ch) + shortened
		}

		if len(shortened)+pos <= max {
			shortened = original[0:pos-1] + shortened
			break
		}
	}

	return shortened
}
