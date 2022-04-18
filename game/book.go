package game

import (
	"bytes"
	"math"
	"strings"
)

const readPDFNewLine string = "📒📒\n"

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

	printBuf := bytes.NewBuffer(nil)

	for i := 1; i <= numPages; i++ {
		pg := b.pdf.Page(i)
		for _, t := range pg.Content().Text {
			y = int(math.Round(t.Y))
			if y != py {
				py = y
				buf.WriteString(readPDFNewLine)
				printBuf.WriteString(readPDFNewLine)
				// fmt.Printf("newline: %v\n", readPDFNewLine)

				if t.S != "\n" {
					// fmt.Printf("is newline!------------------------------\n")
					buf.WriteString("\n")
					printBuf.WriteString("\n")
				}
			}
			buf.WriteString(t.S)
			printBuf.WriteString(t.S)
			// x := int(math.Round(t.X))
			// fmt.Printf("S: %v\t\tC: %v, %v\n", t.S, x, y)
		}
	}

	// sep := "=============================="
	// fmt.Printf("book:\n\n%v\n\n%v\n\n%v\n\n", sep, printBuf.String(), sep)

	c := buf.String()
	c = strings.ReplaceAll(c, readPDFNewLine, "")
	c = strings.ReplaceAll(c, "\n\n", "\n")

	b.buf = bytes.NewBufferString(c)
	return nil
}

// Contents ...
func (b *Book) Contents() string {
	// c := b.buf.String()
	// c = strings.ReplaceAll(c, readPDFNewLine, "")
	// c = strings.ReplaceAll(c, "\n\n", "\n")

	return b.buf.String()
}
