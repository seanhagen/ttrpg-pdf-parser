/*
Copyright © 2022 Sean Patrick Hagen <sean.hagen@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package main

import "github.com/seanhagen/numenera/cmd"

func main() {
	cmd.Execute()
}

/*
var pdfPathFlag = flag.String("pdf", "", "path to the PDF file to parse")
var blankoutsPathFlag = flag.String("blankouts", "", "path to file containing lines to remove from PDF text")
var sectionsPathFlag = flag.String("boundaries", "", "path to file containing section boundaries")
var sectionFixesPathFlag = flag.String("section-fixes", "", "path to file containing section fixes")

func init() {
	flag.Parse()
}

func main() {
	if pdfPathFlag == nil || *pdfPathFlag == "" {
		fmt.Printf("use -pdf to specify the PDF to parse!\n")
		flag.Usage()
		os.Exit(1)
	}

	if blankoutsPathFlag == nil || *blankoutsPathFlag == "" {
		fmt.Printf("use -blankouts to specify the path to a file containing lines to remove")
		flag.Usage()
		os.Exit(1)
	}

	if sectionsPathFlag == nil || *sectionsPathFlag == "" {
		fmt.Printf("use -boundaries to specify the path to a file containing section boundaries")
		flag.Usage()
		os.Exit(1)
	}

	if sectionFixesPathFlag == nil || *sectionFixesPathFlag == "" {
		fmt.Printf("use -section-fixes to specify the path to a file containing section fixes")
		flag.Usage()
		os.Exit(1)
	}

	book, err := prepareBook(*pdfPathFlag, *blankoutsPathFlag, *sectionsPathFlag, *sectionFixesPathFlag)
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

func prepareBook(pdfPath, blankoutPath, sectionPath, sectionFixPath string) (*pdf.Book, error) {
	book, err := pdf.OpenBook(pdfPath)
	if err != nil {
		return nil, fmt.Errorf("unable to open book '%v', error: %w", pdfPath, err)
	}

	err = book.LoadBlankoutsFromFile(blankoutPath)
	if err != nil {
		return nil, fmt.Errorf("unable to load blankouts: %w", err)
	}

	err = book.LoadSectionBoundaries(sectionPath)
	if err != nil {
		return nil, fmt.Errorf("unable to load boundaries: %w", err)
	}

	err = book.LoadSectionFixes(sectionFixPath)
	if err != nil {
		return nil, fmt.Errorf("unable to load section fixes: %w", err)
	}

	err = book.Read()
	if err != nil {
		return nil, fmt.Errorf("unable to read PDF: %w", err)
	}

	err = book.ParseSections()
	if err != nil {
		return nil, fmt.Errorf("unable to parse sections: %w", err)
	}

	return book, nil
}

*/
