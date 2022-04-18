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
package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/seanhagen/ttrpg-pdf-parser/pdf"
	"github.com/spf13/cobra"
)

var (
	pdfPath       string
	blankoutsPath string
	sectionsPath  string
	fixesPath     string

	output io.Writer

	outputFilePath string
	outputFile     *os.File

	outputToStdout bool
	outputToFile   bool

	book *pdf.Book
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tpp",
	Short: "TTRPG PDF Parser: Pull items & creatures out of TTRPG PDFs",
	// 	Long: `A longer description that spans multiple lines and likely contains
	// examples and usage of using your application. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },

	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		writers := []io.Writer{}

		if outputToFile {
			f, err := os.OpenFile(outputFilePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
			if err != nil {
				return fmt.Errorf("unable to open output file: %w", err)
			}
			outputFile = f
			writers = append(writers, f)
		}

		if outputToStdout {
			writers = append(writers, os.Stdout)
		}

		output = io.MultiWriter(writers...)

		var err error
		book, err = pdf.OpenBook(pdfPath)
		if err != nil {
			return fmt.Errorf("unable to open book '%v', error: %w", pdfPath, err)
		}

		err = book.LoadBlankoutsFromFile(blankoutsPath)
		if err != nil {
			return fmt.Errorf("unable to load blankouts: %w", err)
		}

		err = book.LoadSectionBoundaries(sectionsPath)
		if err != nil {
			return fmt.Errorf("unable to load boundaries: %w", err)
		}

		err = book.LoadSectionFixes(fixesPath)
		if err != nil {
			return fmt.Errorf("unable to load section fixes: %w", err)
		}

		err = book.Read()
		if err != nil {
			return fmt.Errorf("unable to read PDF: %w", err)
		}

		err = book.ParseSections()
		if err != nil {
			return fmt.Errorf("unable to parse sections: %w", err)
		}

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().
		StringVar(&pdfPath, "pdf", "", "path to PDF file to parse")

	rootCmd.PersistentFlags().
		StringVar(&blankoutsPath, "blankouts", "", "path to TXT file containing lines to remove from PDF text")

	rootCmd.PersistentFlags().
		StringVar(&sectionsPath, "sections", "", "path to JSON file containing section boundaries")

	rootCmd.PersistentFlags().
		StringVar(&fixesPath, "fixes", "", "path to file containing section fixes")

	rootCmd.PersistentFlags().
		BoolVar(&outputToStdout, "stdout", true, "should the resources be printed to STDOUT (default: yes)")

	rootCmd.PersistentFlags().
		BoolVar(&outputToFile, "file", false, "should the resources be written to a file (default: no)")

	rootCmd.PersistentFlags().
		StringVar(&outputFilePath, "file-path", "./output.txt", "path to a file to write the resources to (default: output.txt)")

	// // Cobra also supports local flags, which will only run
	// // when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
