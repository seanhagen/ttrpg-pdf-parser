package numenera

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItems_OdditySplitLines(t *testing.T) {
	input := loadStringFromFile(t, "oddity-split-input.txt")
	var expect []string
	loadDataFromFile(t, "oddity-split-expect.json", &expect)

	var blankouts []string
	loadDataFromFile(t, "new-artifacts-blankouts.json", &blankouts)

	got := SplitOddityText(input, blankouts)

	assert.Equal(t, expect, got)
}

func TestItems_NewOddityFromLine(t *testing.T) {
	tests := []struct {
		Input  string
		Expect *Oddity
	}{}

	loadDataFromFile(t, "new-oddities.json", &tests)

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			got := NewOddity(tt.Input, []string{})
			assert.Equal(t, tt.Expect, got)
		})
	}
}
