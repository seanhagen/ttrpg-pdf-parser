package items

import (
	"regexp"
	"strings"
)

var (
	pageNumberRE = regexp.MustCompile(`^(\d+)(\w+)`)
)

var replaceWithNothing = []string{
	`“Completed a grafting procedure today with an entirely unknown piece of the numenera. [NAME REDACTED] called it a calridian proboscis. Its wrappings had the mark of the alizarin market. This would lead one to believe that calridian proboscides are neither crafted nor found elements of the numenera, but rather organs from living, breathing creatures. If there is such a thing as a calridian out there, I for one have certainly never seen it. The operation itself was not a success, per se. But I would surely like to get my hands on another one of those devices.” ~notes from Tireso Hinyer, Draolis chiurgeon`,
	`If a singularity detonation is placed in an annihilation chamber, the resulting explosion utterly destroys everything within long range.Singularity detonation page 285 104`,
}

var (
	cypherHeadingReplacements = []string{
		// ".", ".\n",
		"Level:", "\n",
		"Internal:", "\n📒Internal📒 ",
		"Wearable:", "\n📒Wearable📒 ",
		"Usable:", "\n📒Usable📒 ",
		"Effect:", "\n📒Effect📒 ",
	}
	artifactHeadingReplacements = []string{
		"Level: ", "\nLevel: ",
		"Form: ", "\nForm: ",
		"Effect: ", "\nEffect: ",
		"Depletion: ", "\nDepletion: ",
	}
)

var (
	cypherReplacer   = strings.NewReplacer(cypherHeadingReplacements...)
	artifactReplacer = strings.NewReplacer(artifactHeadingReplacements...)
)

var (
	artifactSplitterRE = regexp.MustCompile(`Depletion: (—)|([\d])(–\d)? in (\d+)?(d)?(\d+)(\s?[\+\-]\s?\d+)?( .*?\(.*?\))?`)
	cypherSplitterRE   = regexp.MustCompile(`([^\s.][^.]*)?(Level:) (\d+)?(d)?(\d+)(\s?[\+\-]\s?\d+)?`)
)

// SplitArtifactText ...
func SplitArtifactText(input string) []string {
	input = artifactSplitterRE.ReplaceAllString(input, " Depletion: $1 | $2$3 in $4$5$6$7 $8\n")
	// input = strings.ReplaceAll(input, "Depletion: Depletion:", "Depletion: ")
	input = strings.ReplaceAll(input, "Depletion:  | ", "Depletion: ")
	input = strings.ReplaceAll(input, ".Depl", ". Depl")
	input = strings.ReplaceAll(input, " |  in", "")

	for _, v := range replaceWithNothing {
		input = strings.ReplaceAll(input, v, " ")
	}

	bits := strings.Split(input, "\n")
	var output []string
	for _, v := range bits {
		// fmt.Printf("vvvvv: %v\n", v)
		v = pageNumberRE.ReplaceAllString(v, "$2 ")
		v = strings.ReplaceAll(v, "  ", " ")
		v = strings.ReplaceAll(v, "Depletion: Depletion:", "Depletion:")
		if v != "" {
			output = append(output, strings.TrimSpace(v))
		}
	}

	return output
}

// SplitCypherText takes a whole section of text containing cyphers and
// splits it up into a block of text for each Cypher that can be
// passed into NewCypher
func SplitCypherText(input string) []string {
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
