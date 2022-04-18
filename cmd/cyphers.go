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
	cypherTemplatePath string

	//go:embed templates/cyphers.txt.tmpl
	defaultCypherTemplate string

	cypherTemplate *template.Template
)

// cyphersCmd represents the cyphers command
var cyphersCmd = &cobra.Command{
	Use:   "cyphers",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		t := template.New("cypherDisplay")

		if cypherTemplatePath != "" {
			t, err = t.ParseFiles(cypherTemplatePath)
		} else {
			t, err = t.Parse(defaultCypherTemplate)
		}

		if err != nil {
			return fmt.Errorf("unable to parse cypher template: %w", err)
		}

		cypherTemplate = t

		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		cypherTxt := book.GetSection("CYPHERS")
		splitCypherLines := items.SplitCypherText(cypherTxt, book.GetBlankouts())

		var cyphers []*items.Cypher
		for _, l := range splitCypherLines {
			c := items.NewCypher(l)
			if c != nil {
				cyphers = append(cyphers, c)
			}
		}

		for _, c := range cyphers {
			err := cypherTemplate.Execute(output, c)
			if err != nil {
				return fmt.Errorf("unable to output cypher: %w", err)
			}
		}

		return nil
	},
}

func init() {
	numeneraCmd.AddCommand(cyphersCmd)

	cyphersCmd.PersistentFlags().
		StringVar(&cypherTemplatePath, "templates", "",
			"path to template to use when rendering cyphers, defaults to built-in")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cyphersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cyphersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
