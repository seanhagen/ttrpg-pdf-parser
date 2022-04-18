package game

import (
	"bytes"
	"os"

	"github.com/ledongthuc/pdf"
)

// Book ...
type Book struct {
	file *os.File
	pdf  *pdf.Reader

	buf *bytes.Buffer

	sectionBoundaries sectionList
	sections          map[string]string
	sectionFixes      SectionFixList

	blankouts []string
}

// OpenBook ...
func OpenBook(path string) (*Book, error) {
	f, r, err := pdf.Open(path)
	if err != nil {
		return nil, err
	}

	b := &Book{
		file:              f,
		pdf:               r,
		sectionBoundaries: sectionList{},
		sections:          map[string]string{},
		blankouts:         []string{},
	}

	return b, nil
}
