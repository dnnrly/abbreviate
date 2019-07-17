/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/dnnrly/abbreviate/domain"
	"github.com/spf13/cobra"
)

// pascalCmd represents the pascal command
var pascalCmd = &cobra.Command{
	Use:   "pascal [string]",
	Short: "Abbreviate a string and convert it to pascal case",
	Long:  `Abbreviate a string and convert it to pascal case.`,
	Args:  validateArgPresent,
	Run: func(cmd *cobra.Command, args []string) {
		abbr := domain.AsPascal(matcher, args[0], optMax)

		fmt.Printf("%s", abbr)
		if optNewline {
			fmt.Printf("\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(pascalCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pascalCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pascalCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
