package items

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItems_ArtifactSplitLinesFull(t *testing.T) {
	input := loadStringFromFile(t, "artifacts-split-input.txt")
	var expect []string
	loadDataFromFile(t, "artifacts-split-expect.json", &expect)

	var blankouts []string
	loadDataFromFile(t, "new-artifacts-blankouts.json", &blankouts)
	got := SplitArtifactText(input, blankouts)

	assert.Equal(t, expect, got)
}

func TestItems_NewArtifact(t *testing.T) {
	tests := []struct {
		Input  string
		Expect *Artifact
	}{}

	loadDataFromFile(t, "new-artifacts.json", &tests)

	var blankouts []string
	loadDataFromFile(t, "new-artifacts-blankouts.json", &blankouts)

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i+1), func(t *testing.T) {
			got := NewAritfact(tt.Input, blankouts)
			assert.Equal(t, tt.Expect, got)
		})
	}
}
