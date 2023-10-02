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
	"sort"
	"strings"

	"github.com/spf13/cobra"

	"github.com/dnnrly/abbreviate/data"
	"github.com/dnnrly/abbreviate/domain"
)

var (
	optList            = false
	optNewline         = true
	optLanguage        = "en-us"
	optSet             = "common"
	optCustom          = ""
	optMax             = 0
	optFrmFront        = false
	optStrategy        = "lookup"
	optRemoveStopwords = false

	matcher     *domain.Matcher
	abbreviator domain.Abbreviator
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "abbreviate [action]",
	Short: "Shorten your string using common abbreviations",
	Long: `This tool will attempt to shorten the string provided using common abbreviations
specified by language and 'set'. Word boundaries will be detected using title case
and non-letters.

Hosted on Github - https://github.com/dnnrly/abbreviate

I'm really interested in how you feel about this tool. Please take the time to fill
in this short survey:
https://forms.gle/6xV1gB8yKGdmuHJ78

If you spot a bug, feel free to raise an issue or fix it and make a pull
request. We're really interested to see more abbreviations added or corrected.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if optList {
			fmt.Printf("Available languages and abbreviation sets:\n")
			languages := []string{}
			for l := range data.Abbreviations {
				languages = append(languages, l)
			}
			sort.Strings(languages)

			for _, l := range languages {
				sets := []string{}
				for s := range data.Abbreviations[l] {
					sets = append(sets, s)
				}

				sort.Strings(sets)
				for _, s := range sets {
					fmt.Printf("--language %s --set %s\n", l, s)
				}
			}
			os.Exit(0)
		}

		switch optStrategy {
		case "lookup":
			matcher = setMatcher()
			abbreviator = matcher
		case "no-abbreviation":
			abbreviator = &domain.NoAbbreviator{}
		default:
			fmt.Fprintf(
				os.Stderr,
				`Error: unknown abbreviation strategy '%s'. Only allowed %s
	Run 'abbreviate --help' for usage.
	`,
				optStrategy, strings.Join([]string{"lookup", "no-abbreviation"}, ", "))
			os.Exit(1)
		}
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
	rootCmd.PersistentFlags().BoolVarP(&optFrmFront, "from-front", "", optFrmFront, "Shorten from the front")
	rootCmd.PersistentFlags().StringVarP(&optStrategy, "strategy", "", optStrategy, "Abbreviation strategy")
	rootCmd.PersistentFlags().BoolVarP(&optRemoveStopwords, "no-stopwords", "", optRemoveStopwords, "Remove stopwords from abbreviation")
}

func setMatcher() *domain.Matcher {
	if optCustom == "" {
		return data.Abbreviations[optLanguage][optSet]
	} else {
		all := ""
		buf, err := os.ReadFile(optCustom)
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"Unable to open custom abbreviations file: %s\n",
				err,
			)
			os.Exit(1)
		}

		all = string(buf)
		return domain.NewMatcherFromString(all)
	}
}

func validateArgPresent(cmd *cobra.Command, args []string) error {
	if !optList && len(args) != 1 {
		return errors.New("requires a string to abbreviate")
	}

	return nil
}
