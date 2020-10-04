package data

import "github.com/dnnrly/abbreviate/domain"

// Sets represents a list named of sets of abbreviations
type Sets map[string]*domain.Matcher

// LangageSets is the collection of Set for a particular language
type LangageSets map[string]Sets

// Abbreviations is the list of all abbreviations by language then set
var Abbreviations = LangageSets{
	"en-us": Sets{
		"common": domain.NewMatcher(enUSCommonMainWords, enUSCommonPrefixes),
	},
}
