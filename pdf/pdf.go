package pdf

import (
	"bytes"
	"math"
	"os"

	"github.com/ledongthuc/pdf"
)

const readPDFNewLine string = "📒📒\n"

// Book ...
type Book struct {
	file *os.File
	pdf  *pdf.Reader

	buf *bytes.Buffer

	sectionBoundaries sectionList
	sections          map[string]string
	blankouts         []string
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

// Close  ...
func (b *Book) Close() error {
	if b.file == nil {
		return nil
	}

	b.buf = nil
	b.sectionBoundaries = nil
	b.sections = nil
	b.blankouts = nil

	err := b.file.Close()
	b.file = nil

	return err
}

// Read ...
func (b *Book) Read() error {
	numPages := b.pdf.NumPage()
	buf := bytes.NewBuffer(nil)
	y := int(math.Round(b.pdf.Page(1).Content().Text[0].Y))
	py := y

	for i := 1; i <= numPages; i++ {
		pg := b.pdf.Page(i)
		for _, t := range pg.Content().Text {
			y = int(math.Round(t.Y))
			if y != py {
				py = y
				buf.WriteString(readPDFNewLine)
				// buf.WriteString("\n")
			}
			buf.WriteString(t.S)
			// if t.S != "\n" {
			// }
		}
	}

	b.buf = buf
	return nil
}
