package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/seanhagen/numenera/items"
	"github.com/seanhagen/numenera/pdf"
)

var pdfPathFlag = flag.String("pdf", "", "path to the PDF file to parse")

var fixes = pdf.SectionFixList{
	{
		Match: regexp.MustCompile(`([^\s.].*?)?📒+`),
		Fix:   "\n\n$1⏱\n",
	},
	{
		Match: regexp.MustCompile(`[^.]?📒\s*`),
		Fix:   " ",
	},
	{
		Match: regexp.MustCompile(`\s+⏱\s*`),
		Fix:   " ",
	},
	{
		Match: regexp.MustCompile(`⏱\s*`),
		Fix:   "\n",
	},
}

func init() {
	flag.Parse()
}

func main() {
	if pdfPathFlag == nil || *pdfPathFlag == "" {
		fmt.Printf("use -pdf to specify the PDF to parse!\n")
		os.Exit(1)
	}

	book, err := prepareBook(*pdfPathFlag)
	if err != nil {
		fmt.Printf("Unable to get book '%v', error: %v\n", *pdfPathFlag, err)
		os.Exit(1)
	}

	cypherTxt := book.GetSection("CYPHERS")
	artifactTxt := book.GetSection("ARTIFACTS")

	splitCypherLines := items.SplitCypherText(cypherTxt, book.GetBlankouts())
	splitArtifactLines := items.SplitArtifactText(artifactTxt, book.GetBlankouts())

	var cyphers []*items.Cypher
	var artifacts []*items.Artifact

	for _, l := range splitCypherLines {
		c := items.NewCypher(l)
		if c != nil {
			cyphers = append(cyphers, c)
		}
	}

	for _, l := range splitArtifactLines {
		a := items.NewAritfact(l, book.GetBlankouts())
		if a != nil {
			artifacts = append(artifacts, a)
		}
	}

	fmt.Printf("got %v cyphers, and %v artifacts!\n", len(cyphers), len(artifacts))

	for _, c := range cyphers {
		fmt.Printf("CYPHER: %v ( Level: %v )\n------------------------------\n", c.Name, c.Level)

		if c.Internal != "" {
			fmt.Printf("Internal: %v\n", c.Internal)
		}
		if c.Usable != "" {
			fmt.Printf("Usable: %v\n", c.Usable)
		}
		if c.Wearable != "" {
			fmt.Printf("Wearable: %v\n", c.Wearable)
		}

		fmt.Printf("--------------------\n")
		fmt.Printf("Effect:\n%v\n\n\n", c.Effect)
	}

	for _, a := range artifacts {
		fmt.Printf("ARTIFACT: %v ( Level %v)\n------------------------------\n", a.Name, a.Level)
		fmt.Printf("Form: %v\nDepletion: %v\nEffect:\n%v\n\n\n", a.Form, a.Depletion, a.Effect)
	}
}

func prepareBook(path string) (*pdf.Book, error) {
	book, err := pdf.OpenBook(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open book '%v', error: %w", path, err)
	}

	err = book.LoadBlankoutsFromFile("./blankouts.txt")
	if err != nil {
		return nil, fmt.Errorf("unable to load blankouts: %w", err)
	}

	err = book.LoadSectionBoundaries("./boundaries.json")
	if err != nil {
		return nil, fmt.Errorf("Unable to load boundaries: %w", err)
	}

	err = book.Read()
	if err != nil {
		return nil, fmt.Errorf("unable to read PDF: %w", err)
	}

	err = book.ParseSections(fixes)
	if err != nil {
		return nil, fmt.Errorf("unable to parse sections: %w", err)
	}

	return book, nil
}
