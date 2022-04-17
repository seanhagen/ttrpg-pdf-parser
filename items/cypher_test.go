package items

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumenera_CypherSplitLines(t *testing.T) {
	input := loadStringFromFile(t, "cypher-split-input.txt")
	var expect []string
	loadDataFromFile(t, "cypher-split-expect.json", &expect)

	got := SplitCypherText(input)

	assert.Equal(t, expect, got)
}

func TestNumenera_Cypher(t *testing.T) {
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
