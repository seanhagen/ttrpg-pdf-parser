package numenera

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItems_CypherSplitLines(t *testing.T) {
	input := loadStringFromFile(t, "cypher-split-input.txt")
	var expect []string
	loadDataFromFile(t, "cypher-split-expect.json", &expect)

	var blankouts []string
	loadDataFromFile(t, "new-artifacts-blankouts.json", &blankouts)

	got := SplitCypherText(input, blankouts)

	assert.Equal(t, expect, got)
}

func TestItems_NewCypherFromLine(t *testing.T) {
	tests := []struct {
		Input  string
		Expect *Cypher
	}{}

	loadDataFromFile(t, "new-cyphers.json", &tests)

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			got := NewCypher(tt.Input)
			assert.Equal(t, tt.Expect, got)
		})
	}
}
