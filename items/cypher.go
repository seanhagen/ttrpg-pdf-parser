package items

import (
	"strings"
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
	parts := strings.Split(cypherReplacer.Replace(line), "\n")

	c := &Cypher{
		Name:  strings.TrimSpace(parts[0]),
		Level: strings.TrimSpace(parts[1]),
	}

	for _, v := range parts[2:] {
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
