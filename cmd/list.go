package cmd

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	Run:  ListRun,
	Args: cobra.NoArgs,
}

func ListRun(cmd *cobra.Command, _ []string) {
	var workspaces = viper.GetStringMapStringSlice("workspaces")
	var workspaceList = make([]string, 0, len(workspaces))
	for key := range workspaces {
		workspaceList = append(workspaceList, key)
	}
	sort.Strings(workspaceList)
	for _, key := range workspaceList {
		currentWorkspace := viper.GetString("settings.currentworkspace")
		boringFlag, err := cmd.Flags().GetBool("boring")
		if err == nil && !boringFlag {
			if key == currentWorkspace {
				fmt.Print(" ✔︎ ")
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
