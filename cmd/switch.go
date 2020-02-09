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
	"os"
	"os/exec"
	"strings"
)

// switchCmd represents the switch command
var switchCmd = &cobra.Command{ // TODO: metadata
	Use:   "switch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: SwitchRun,
	Args: cobra.ExactValidArgs(1),
	ValidArgs: SwitchValidArgs(),
}

func SwitchValidArgs() []string {
	initConfig()
	var workspaces = viper.GetStringMapStringSlice("workspaces")
	var keys = make([]string, 0, len(workspaces))
	for k := range workspaces {
        keys = append(keys, k)
	}
	return keys
}

func SwitchRun(_ *cobra.Command, args []string) {
	var requestedWorkspace = args[0]
	var currentWorkspace = viper.GetString("settings.currentworkspace")

	if requestedWorkspace == currentWorkspace {
		fmt.Printf("Current workspace is already %s\n", requestedWorkspace)
		return
	}

	fmt.Printf("Switching from %s to %s\n", currentWorkspace, requestedWorkspace)

	// u:beforeDown, o:down, u:afterDown, u:beforeUp, n:up, u:afterUp
	execOrder := []string{
		"hooks.beforedown",
		fmt.Sprintf("workspaces.%s.down", currentWorkspace),
		"hooks.afterdown",
		"hooks.beforeup",
		fmt.Sprintf("workspaces.%s.up", requestedWorkspace),
		"hooks.afterup",
	}

	for _, cfgPath := range execOrder {
		commands := viper.GetStringSlice(cfgPath)
		for _, cmd := range commands {
			cmd_array := strings.Split(cmd, " ") // TODO: do this more intelligently, or use an array in the YAML
			thisCommand := exec.Command(cmd_array[0], cmd_array[1:]...)
			thisCommand.Env = os.Environ()
			thisCommand.Stdout = os.Stdout
			thisCommand.Stderr = os.Stderr
			_ = thisCommand.Run()
		}
	}

	viper.Set("settings.currentworkspace", requestedWorkspace)
	_ = viper.WriteConfig()

}

func init() {
	rootCmd.AddCommand(switchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// switchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// switchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
