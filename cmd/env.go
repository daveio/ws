package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// envCmd represents the env command
var envCmd = &cobra.Command{
	Use:     "env",
	Example: `ws env`,
	Short:   "Output environment for current workspace for your shell to eval",
	Long: `Outputs the environment values for the current workspace. Used
by shell integration, to be eval'd or sourced at each prompt render.`,
	Run:  EnvRun,
	Args: cobra.NoArgs, // cobra.ExactValidArgs(1),
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
