package items

import (
	"strings"
)

// Artifact ...
type Artifact struct {
	Name      string
	Level     string
	Form      string
	Effect    string
	Depletion string
}

// NewAritfact ...
func NewAritfact(line string) *Artifact {
	a := &Artifact{}

	parts := strings.Split(artifactReplacer.Replace(line), "\n")

	for i, v := range parts {
		v = strings.TrimSpace(v)
		if i == 0 {
			a.Name = strings.TrimSpace(v)
			continue
		}

		if strings.Contains(v, "Level: ") {
			a.Level = strings.ReplaceAll(v, "Level: ", "")
			continue
		}

		if strings.Contains(v, "Form: ") {
			a.Form = strings.ReplaceAll(v, "Form: ", "")
			continue
		}

		if strings.Contains(v, "Effect: ") {
			a.Effect = strings.ReplaceAll(v, "Effect: ", "")
			continue
		}

		if strings.Contains(v, "Depletion: ") {
			a.Depletion = strings.ReplaceAll(v, "Depletion: ", "")
			continue
		}
	}

	return a
}
