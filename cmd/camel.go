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
	"strings"

	"github.com/dnnrly/abbreviate/domain"
	"github.com/spf13/cobra"
)

// camelCmd represents the camel command
var camelCmd = &cobra.Command{
	Use:   "camel [string]",
	Short: "Abbreviate a string and convert it to camel case",
	Long:  `Abbreviate a string and convert it to camel case.`,
	Args:  validateArgPresent,
	Run: func(cmd *cobra.Command, args []string) {
		abbr := domain.AsPascal(matcher, args[0], optMax)

		ch := string(abbr[0])

		fmt.Printf("%s%s", strings.ToLower(ch), abbr[1:])
		if optNewline {
			fmt.Printf("\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(camelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// camelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// camelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
