package numenera

import (
	"regexp"
)

var (
	pageNumberRE = regexp.MustCompile(`^(\d+)(\w+)`)
	alphaOnlyRE  = regexp.MustCompile(`(\D+)`)
	fixNumberRE  = regexp.MustCompile(`\((\d*)\.\s?(\d* \w)\)`)
	fixSpacesRE  = regexp.MustCompile(`(\S)\s+(\S)`)
)

const (
	pageNumberReplaceWith = ""
	fixNumberReplaceWith  = "($1.$2)"
	fixSpacesWith         = "$1 $2"
)

func fixSpaces(s string) string {
	return fixSpacesRE.ReplaceAllString(s, fixSpacesWith)
}
