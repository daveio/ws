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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

// envCmd represents the env command
var envCmd = &cobra.Command{
	Use: "env",
	Example: `ws env`,
	Short: "Output environment for current workspace for your shell to eval",
	Long: `Outputs the environment values for the current workspace. Used
by shell integration, to be eval'd or sourced at each prompt render.`,
	Run:       EnvRun,
	Args:      cobra.NoArgs, // cobra.ExactValidArgs(1),
	// ValidArgs: []string{"bash", "zsh", "fish"},
}

func EnvRun(_ *cobra.Command, _ []string) {
	if viper.IsSet("settings.currentworkspace") &&
		viper.IsSet(fmt.Sprintf("workspaces.%s.env", viper.GetString("settings.currentworkspace"))) {
		var currentWorkspace = viper.GetString("settings.currentworkspace")
		var environment = viper.GetStringMapString(fmt.Sprintf("workspaces.%s.env", currentWorkspace))
		for k := range environment {
			fmt.Printf("%s=%s\n", strings.ToUpper(k), environment[k])
		}
		fmt.Printf("WS_ENV=%s\n", currentWorkspace)
	}
}

func init() {
	rootCmd.AddCommand(envCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// envCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// envCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
