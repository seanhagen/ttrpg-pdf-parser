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
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var (
	printOutputPath string
)

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print <input> <output>",
	Short: "Outputs the text of the PDF into a normal text file",
	Long: `Print reads all the text from the provided PDF, and outputs
all the text found within into a .txt file.

This is handy when trying to build blankouts or section boundaries,
as the text that tpp is able to read isn't necessarily what you see
when you look a PDF!`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return cmd.Help()
		}
		printOutputPath = args[1]

		if printOutputPath == "" {
			return fmt.Errorf("")
		}
		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		c := bytes.NewBufferString(book.Contents())

		of, err := os.OpenFile(printOutputPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
		if err != nil {
			return fmt.Errorf("unable to open output file '%v', error: %w", printOutputPath, err)
		}

		_, err = io.Copy(of, c)
		if err != nil {
			return fmt.Errorf("unable to write to file: %w", err)
		}

		fmt.Println("copy complete!")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(printCmd)
}
