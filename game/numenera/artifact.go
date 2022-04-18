package numenera

import (
	"regexp"
	"strings"
)

var (
	artifactHeadingReplacements = []string{
		"Level: ", "\nLevel: ",
		"Form: ", "\nForm: ",
		"Effect: ", "\nEffect: ",
		"Depletion: ", "\nDepletion: ",
	}

	artifactReplacer = strings.NewReplacer(artifactHeadingReplacements...)

	artifactSplitterREs = []*regexp.Regexp{
		regexp.MustCompile(`Depletion: (—)\s`),
		regexp.MustCompile(`Depletion: ([\d])(–\d)? in (\d+)?d?(\d+)(\s?[\+\-]\s?\d+)?(.*?)Level`),
		regexp.MustCompile(`⏱(.*)\)\s(.*)⏱`),
		regexp.MustCompile(`⏱\s(.*)⏱\n`),
	}
	artifactSplitterReplacements = []string{
		"Depletion: $1\n\n",
		"Depletion: $1$2 in $3 🎲 $4$5 ⏱$6⏱\nLevel",
		"$1)\n\n$2",
		"\n\n$1",
	}
)

// Artifact ...
type Artifact struct {
	Name      string
	Level     string
	Form      string
	Effect    string
	Depletion string
}

// NewArtifact ...
func NewArtifact(line string, blankouts []string) *Artifact {
	a := &Artifact{}

	line = strings.ReplaceAll(line, "\n", " ")
	toTrim := strings.Split(line, ".")
	for i, v := range toTrim {
		toTrim[i] = strings.TrimSpace(v)
	}
	line = strings.Join(toTrim, ". ")

	parts := strings.Split(artifactReplacer.Replace(line), "\n")

	for i, v := range parts {
		v = fixNumberRE.ReplaceAllString(v, fixNumberReplaceWith)
		v = strings.TrimSpace(v)
		if i == 0 {
			a.Name = strings.TrimSpace(alphaOnlyRE.FindString(parts[0]))
			continue
		}

		if strings.Contains(v, "Level: ") {
			a.Level = applyBlankouts(v, append([]string{"Level: "}, blankouts...))
			continue
		}

		if strings.Contains(v, "Form: ") {
			a.Form = applyBlankouts(v, append([]string{"Form: "}, blankouts...))
			continue
		}

		if strings.Contains(v, "Effect: ") {
			a.Effect = applyBlankouts(v, append([]string{"Effect: "}, blankouts...))
			continue
		}

		if strings.Contains(v, "Depletion: ") {
			a.Depletion = applyBlankouts(v, append([]string{"Depletion: "}, blankouts...))
			continue
		}
	}

	return a
}

func applyBlankouts(s string, b []string) string {
	for _, v := range b {
		s = strings.ReplaceAll(s, v, "")
	}
	return s
}

// SplitArtifactText ...
func SplitArtifactText(input string, blankouts []string) []string {
	for _, b := range blankouts {
		input = strings.ReplaceAll(input, b, "")
	}

	input = strings.ReplaceAll(input, "\n", " ")

	for i, re := range artifactSplitterREs {
		input = re.ReplaceAllString(input, artifactSplitterReplacements[i])
	}

	input = strings.ReplaceAll(input, " 🎲 ", "d")
	input = strings.ReplaceAll(input, ".  ", ". ")

	input = strings.ReplaceAll(input, "Depletion:—", "Depletion: —")
	input = strings.ReplaceAll(input, "Depletion: | ", "Depletion: ")
	input = strings.ReplaceAll(input, ".Depl", ". Depl")
	input = strings.ReplaceAll(input, " |  in", "")

	bits := strings.Split(input, "\n\n")
	var output []string
	for _, v := range bits {
		v = strings.ReplaceAll(v, "\n", " ")
		v = fixSpacesRE.ReplaceAllString(v, fixSpacesWith)
		v = pageNumberRE.ReplaceAllString(v, pageNumberReplaceWith)
		v = strings.ReplaceAll(v, "  ", " ")
		v = strings.TrimSpace(v)
		if v != "" {
			output = append(output, strings.TrimSpace(v))
		}
	}

	return output
}
