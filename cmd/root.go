// Copyright © 2018 Pascal Dennerly
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/cobra"

	"github.com/dnnrly/abbreviate/domain"
)

var (
	optList     = false
	optNewline  = true
	optLanguage = "en-us"
	optSet      = "common"
	optCustom   = ""
	optMax      = 0

	data    = packr.New("abbreviate", "../data")
	matcher *domain.Matcher
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "abbreviate [action]",
	Short: "Shorten your string using common abbreviations",
	Long: `This tool will attempt to shorten the string provided using common abbreviations
specified by language and 'set'. Word boundaries will be detected using title case
and non-letters.

Hosted on Github - https://github.com/dnnrly/abbreviate

If you spot a bug, feel free to raise an issue or fix it and make a pull
request. We're really interested to see more abbreviations added or corrected.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if optList {
			fmt.Printf("Available languages and abbreviation sets:\n")
			items := data.List()
			for _, v := range items {
				parts := strings.Split(v, "/")
				fmt.Printf("--language %s --set %s\n", parts[0], parts[1])
			}
			os.Exit(0)
		}

		path := fmt.Sprintf("%s/%s", optLanguage, optSet)
		all, err := data.FindString(path)
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"Unable to find language '%s' with set '%s'.\n",
				optLanguage,
				optSet,
			)
			os.Exit(1)
		}

		matcher = domain.NewMatcherFromString(all)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(
			os.Stderr,
			`Error: use one of the actions to abbreviate a string.
Run 'abbreviate --help' for usage.
`,
		)
		os.Exit(1)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&optList, "list", "", optList, "List all abbreviate sets by language")
	rootCmd.PersistentFlags().BoolVarP(&optNewline, "newline", "n", optNewline, "Add newline to the end of the string")
	rootCmd.PersistentFlags().StringVarP(&optLanguage, "language", "l", optLanguage, "Language to select")
	rootCmd.PersistentFlags().StringVarP(&optSet, "set", "s", optSet, "Abbreviation set")
	rootCmd.PersistentFlags().StringVarP(&optCustom, "custom", "c", optCustom, "Custom abbreviation set")
	rootCmd.PersistentFlags().IntVarP(&optMax, "max", "m", optMax, "Maximum length of string, keep on abbreviating while the string is longer than this limit")
}

func validateArgPresent(cmd *cobra.Command, args []string) error {
	if !optList && len(args) != 1 {
		return errors.New("requires a string to abbreviate")
	}

	return nil
}
