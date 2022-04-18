package numenera

import (
	"regexp"
	"strings"
)

var (
	cypherSplitterRE = regexp.MustCompile(`([^\s.][^.]*)?(Level:) (\d+)?(d)?(\d+)(\s?[\+\-]\s?\d+)?`)

	cypherHeadingReplacements = []string{
		// ".", ".\n",
		"Level:", "\n",
		"Internal:", "\n📒Internal📒 ",
		"Wearable:", "\n📒Wearable📒 ",
		"Usable:", "\n📒Usable📒 ",
		"Effect:", "\n📒Effect📒 ",
	}

	cypherReplacer = strings.NewReplacer(cypherHeadingReplacements...)
)

// Cypher ...
type Cypher struct {
	Name     string
	Level    string
	Wearable string
	Internal string
	Usable   string
	Effect   string
}

// NewCypher ...
func NewCypher(line string) *Cypher {
	line = strings.ReplaceAll(line, "\n", " ")
	toTrim := strings.Split(line, ".")
	for i, v := range toTrim {
		toTrim[i] = strings.TrimSpace(v)
	}
	line = strings.Join(toTrim, ". ")

	tmp := strings.Split(cypherReplacer.Replace(line), "\n")
	var parts []string
	for _, v := range tmp {
		v = strings.TrimSpace(v)
		if v != "" {
			parts = append(parts, v)
		}
	}

	name := strings.TrimSpace(alphaOnlyRE.FindString(parts[0]))

	c := &Cypher{
		Name:  name,
		Level: strings.TrimSpace(parts[1]),
	}

	for _, v := range parts[2:] {
		v = fixNumberRE.ReplaceAllString(v, fixNumberReplaceWith)

		if strings.Contains(v, "📒Internal📒 ") {
			v = strings.ReplaceAll(v, "📒Internal📒 ", "")
			c.Internal = strings.TrimSpace(v)
		}

		if strings.Contains(v, "📒Wearable📒 ") {
			v = strings.ReplaceAll(v, "📒Wearable📒 ", "")
			c.Wearable = strings.TrimSpace(v)
		}

		if strings.Contains(v, "📒Usable📒 ") {
			v = strings.ReplaceAll(v, "📒Usable📒 ", "")
			c.Usable = strings.TrimSpace(v)
		}

		if strings.Contains(v, "📒Effect📒 ") {
			v = strings.ReplaceAll(v, "📒Effect📒 ", "")
			c.Effect = strings.TrimSpace(v)
		}
	}

	return c
}

// SplitCypherText takes a whole section of text containing cyphers and
// splits it up into a block of text for each Cypher that can be
// passed into NewCypher
func SplitCypherText(input string, blankouts []string) []string {
	var output []string

	matches := cypherSplitterRE.FindAllString(input, -1)

	var reversed = make([]string, len(matches))
	j := 0
	for i := len(matches) - 1; i >= 0; i-- {
		reversed[j] = matches[i]
		j++
	}

	for _, v := range reversed {
		parts := strings.Split(input, v)

		if len(parts) == 2 {
			bit := v + parts[1]
			input = strings.Replace(input, bit, "", 1)
			bit = pageNumberRE.ReplaceAllString(bit, "$2")
			bit = cypherSplitterRE.ReplaceAllString(bit, "😆 $1 😆😆 $2 $3$4$5$6 😆😆😆")
			bit = strings.ReplaceAll(bit, "😆😆😆", " ")
			bit = strings.ReplaceAll(bit, "😆😆 ", "")
			bit = strings.ReplaceAll(bit, "😆 ", "")
			bit = strings.ReplaceAll(bit, "  ", " ")
			bit = strings.TrimSpace(bit)
			for _, b := range blankouts {
				bit = strings.ReplaceAll(bit, b, " ")
			}

			output = append(output, bit)
		}
	}

	reversed = make([]string, len(output))
	j = 0
	for i := len(output) - 1; i >= 0; i-- {
		reversed[j] = output[i]
		j++
	}

	return reversed
}
