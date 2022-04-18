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
	_ "embed"
	"fmt"
	"text/template"

	"github.com/seanhagen/ttrpg-pdf-parser/items"
	"github.com/spf13/cobra"
)

var (
	oddityTemplatePath string

	//go:embed templates/oddities.txt.tmpl
	defaultOddityTemplate string

	oddityTemplate *template.Template
)

// odditiesCmd represents the oddities command
var odditiesCmd = &cobra.Command{
	Use:   "oddities",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		t := template.New("oddityDisplay")

		if oddityTemplatePath != "" {
			t, err = t.ParseFiles(oddityTemplatePath)
		} else {
			t, err = t.Parse(defaultOddityTemplate)
		}

		if err != nil {
			return fmt.Errorf("unable to parse oddity template: %w", err)
		}

		oddityTemplate = t

		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		oddityTxt := book.GetSection("ODDITIES")
		splitOddityLines := items.SplitOddityText(oddityTxt, book.GetBlankouts())

		var oddities []*items.Oddity
		for _, l := range splitOddityLines {
			c := items.NewOddity(l, book.GetBlankouts())
			if c != nil {
				oddities = append(oddities, c)
			}
		}

		for _, c := range oddities {
			err := oddityTemplate.Execute(output, c)
			if err != nil {
				return fmt.Errorf("unable to output oddity: %w", err)
			}
		}

		return nil
	},
}

func init() {
	numeneraCmd.AddCommand(odditiesCmd)

	odditiesCmd.PersistentFlags().
		StringVar(&oddityTemplatePath, "templates", "",
			"path to template to use when rendering oddities, defaults to built-in")
}
