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
	"errors"
	"fmt"

	"github.com/dnnrly/abbreviate/domain"
	"github.com/spf13/cobra"
)

var (
	optSeperator = "_"
)

// originalCmd represents the original command
var originalCmd = &cobra.Command{
	Use:   "original [string]",
	Short: "Abbreviate the string using the original word boundary seperators",
	Args: func(cmd *cobra.Command, args []string) error {
		if !optList && len(args) != 1 {
			return errors.New("requires a string to abbreviate")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		abbr := domain.AsOriginal(matcher, args[0], optMax)

		fmt.Printf("%s", abbr)
		if optNewline {
			fmt.Printf("\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(originalCmd)
}
