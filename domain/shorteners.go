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
	for len(original) > 0 && len(original)+len(word)+len(shortened) > max {
		remaining, ch := lastChar(original)
		if unicode.IsLetter(ch) {
			word = string(ch) + word
			if unicode.IsUpper(ch) {
				word = strings.ToLower(word)
				abbr := matcher.Match(word)
				abbr = strings.Title(abbr)
				shortened = abbr + shortened
				word = ""
			}
		} else {
			if word != "" {
				abbr := matcher.Match(word)
				shortened = abbr + shortened
				word = ""
			}
			shortened = string(ch) + shortened
		}
		original = remaining
	}
	if len(original)+len(shortened) > max {
		word = matcher.Match(word)
	}
	shortened = original + word + shortened

	return shortened
}

func lastChar(str string) (string, rune) {
	l := len(str)

	switch l {
	case 0:
		return "", rune(0)

	case 1:
		return "", []rune(str)[0]
	}

	return str[0 : l-1], []rune(str)[l-1:][0]
}
