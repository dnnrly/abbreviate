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

	"github.com/dnnrly/abbreviate/domain"
	"github.com/spf13/cobra"
)

var (
	optSeperator = ""
)
// separatedCmd represents the Separated command
var separatedCmd = &cobra.Command{
	Use:   "separated [string]",
	Short: "Abbreviate a string and convert it using separator passed",
	Long: `Abbreviate a string and convert it using separator passed.

Where a string is not shortened, it will be converted to your case anyway, even
if this means that the string will end up longer.`,
	Args: validateArgPresent,
	Run: func(cmd *cobra.Command, args []string) {
		abbr := domain.AsSeparated(matcher, args[0], optSeperator, optMax, optFrmFront,optStrategy)

		fmt.Printf("%s", abbr)
		if optNewline {
			fmt.Printf("\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(separatedCmd)
	separatedCmd.Flags().StringVarP(&optSeperator, "separator", "S", optSeperator, "Seperator between words and abbreviations when using 'separated' command (default \"\")")
}
