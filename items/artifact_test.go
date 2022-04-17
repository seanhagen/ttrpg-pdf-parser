package items

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArtifacts_SplitLines(t *testing.T) {
	input := loadStringFromFile(t, "artifacts-split-input.txt")
	var expect []string
	loadDataFromFile(t, "artifacts-split-expect.json", &expect)

	got := SplitArtifactText(input)

	assert.Equal(t, expect, got)
}

func TestArtifacts_NewArtifact(t *testing.T) {
	tests := []struct {
		Input  string
		Expect *Artifact
	}{}

	loadDataFromFile(t, "new-artifacts.json", &tests)

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i+1), func(t *testing.T) {
			got := NewAritfact(tt.Input)
			assert.Equal(t, tt.Expect, got)
		})
	}
}
