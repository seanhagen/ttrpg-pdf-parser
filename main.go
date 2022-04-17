package main

import (
	"fmt"
	"os"

	"github.com/seanhagen/numenera/items"
	"github.com/seanhagen/numenera/pdf"
)

const pdfPath = "./out/tech.pdf"

func main() {
	book, err := pdf.OpenBook(pdfPath)
	if err != nil {
		fmt.Printf("unable to open book '%v', error: %v\n", pdfPath, err)
		os.Exit(1)
	}

	err = book.LoadBlankoutsFromFile("./blankouts.txt")
	if err != nil {
		fmt.Printf("Unable to load blankouts: %v\n", err)
		os.Exit(1)
	}

	err = book.LoadSectionBoundaries("./boundaries.json")
	if err != nil {
		fmt.Printf("Unable to load boundaries: %v\n", err)
		os.Exit(1)
	}

	err = book.Read()
	if err != nil {
		fmt.Printf("Unable to read PDF: %v\n", err)
		os.Exit(1)
	}

	err = book.ParseSections()
	if err != nil {
		fmt.Printf("Unable to parse sections: %v\n", err)
		os.Exit(1)
	}

	cypherTxt := book.GetSection("CYPHERS")
	artifactTxt := book.GetSection("ARTIFACTS")

	splitCypherLines := items.SplitCypherText(cypherTxt)
	splitArtifactLines := items.SplitArtifactText(artifactTxt)

	var cyphers []*items.Cypher
	var artifacts []*items.Artifact

	for _, l := range splitCypherLines {
		c := items.NewCypher(l)
		if c != nil {
			cyphers = append(cyphers, c)
		}
	}

	for _, l := range splitArtifactLines {
		a := items.NewAritfact(l)
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

	// fmt.Printf("\n\n\n\n========================================\n\n\n")
	// fmt.Printf("Cyphers: \n%v\n", cypherTxt)
	// fmt.Printf("\n\n\n\n========================================\n\n\n")
	// fmt.Printf("Artifacts: \n%v\n", artifactTxt)
	// fmt.Printf("\n\n\n\n========================================\n\n\n")
	// litter.Config.FieldExclusions = regexp.MustCompile(`^pdf$|^buf$`)
	// litter.Config.HidePrivateFields = false
	// litter.Dump(book)
	//	spew.Dump(book)
}
