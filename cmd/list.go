/*
Copyright © 2020 Dave Williams <dave@dave.io>

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
	"fmt"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{ // TODO: metadata
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: ListRun,
	Args: cobra.NoArgs,
}

func ListRun(cmd *cobra.Command, args []string) {
	var workspaces = viper.GetStringMapStringSlice("workspaces")
	for key := range workspaces {
		currentWorkspace := viper.GetString("settings.currentWorkspace")
		boringFlag, err := cmd.Flags().GetBool("boring")
		if err == nil && !boringFlag {
			if key == currentWorkspace {
				fmt.Print("✅ ")
			} else {
				fmt.Print("   ")
			}
		}
        fmt.Printf("%s\n", key)

	}
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP(
		"boring",
		"b",
		false,
		"Output a plain list without the 'active workspace' indicator")
}
