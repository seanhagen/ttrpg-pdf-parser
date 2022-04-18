package items

import (
	"regexp"
	"strings"
)

var isOddityRE = regexp.MustCompile(`^\d+. `)

var (
	oddityFixes = []*regexp.Regexp{
		regexp.MustCompile(`📒📒\s(\d+)\.\s`),
		regexp.MustCompile(`📒📒\s📒📒.*`),
		regexp.MustCompile(`📒📒\s`),
	}

	oddityReplace = []string{
		"\n\n$1⏱ ",
		"",
		"",
	}
)

// Oddity ...
type Oddity struct {
	Value string
}

// NewOddity ...
func NewOddity(line string, blankouts []string) *Oddity {
	for _, b := range blankouts {
		line = strings.ReplaceAll(line, b, "")
	}

	return &Oddity{line}
}

// SplitOddityText ...
func SplitOddityText(input string, blankouts []string) []string {
	input = applyBlankouts(input, blankouts)
	input = fixSpaces(input)

	for i, f := range oddityFixes {
		input = f.ReplaceAllString(input, oddityReplace[i])
	}

	var lines []string
	for _, v := range strings.Split(input, "\n") {
		v = strings.ReplaceAll(v, "\n", " ")

		if isOddityRE.MatchString(v) {
			v = isOddityRE.ReplaceAllString(v, "")
			v = strings.TrimSpace(v)
			lines = append(lines, v)
		}
	}

	return lines
}
