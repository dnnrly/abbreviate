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
	"fmt"
	"os"
	"strings"

	"github.com/gobuffalo/packr/v2"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	optList     = false
	optLanguage = "en-us"
	optSet      = "common"
	optCustom   = ""
	optMax      = 0

	data = packr.New("abbreviate", "../data")
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "abbreviate [string]",
	Short: "Shorten your string using common abbreviations",
	Long: `This tool will attempt to shorten the string provided using common abbreviations
specified by language and 'set'.

Word boundaries will detect camel case and non-letter`,
	Args: func(cmd *cobra.Command, args []string) error {
		if !optList && len(args) != 1 {
			return errors.New("requires a string to abbreviate")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if optList {
			fmt.Printf("Available languages and abbreviation sets:\n")
			items := data.List()
			for _, v := range items {
				parts := strings.Split(v, "/")
				fmt.Printf("--language %s --set %s\n", parts[0], parts[1])
			}
		}
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
	rootCmd.Flags().BoolVarP(&optList, "list", "", optList, "List all abbreviate sets by language")
	rootCmd.Flags().StringVarP(&optLanguage, "language", "l", optLanguage, "Language to select")
	rootCmd.Flags().StringVarP(&optSet, "set", "s", optSet, "Abbreviation set")
	rootCmd.Flags().StringVarP(&optCustom, "custom", "c", optCustom, "Custom abbreviation set")
	rootCmd.Flags().IntVarP(&optMax, "max", "m", optMax, "Maximum length of string, keep on abbreviating while the string is longer than this limit")
}
