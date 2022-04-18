package items

import (
	"regexp"
	"strings"
)

var isOddityRE = regexp.MustCompile(`^\d+. `)

var (
	oddityFixes = []*regexp.Regexp{
		regexp.MustCompile(`(\d+)\.\s`),
	}

	oddityReplace = []string{
		"\n\n$1. ",
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

	check := strings.Split(input, "\n")
	var lines []string
	for _, v := range check {
		v = strings.ReplaceAll(v, "\n", " ")

		if isOddityRE.MatchString(v) {
			v = isOddityRE.ReplaceAllString(v, "")
			v = strings.TrimSpace(v)
			tmp := strings.Split(v, " ")
			var g []string
			for _, x := range tmp {
				x = strings.TrimSpace(x)
				if x != "" {
					g = append(g, x)
				}
			}
			v = strings.Join(g, " ")
			lines = append(lines, v)
		}
	}

	return lines
}
