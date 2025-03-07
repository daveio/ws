package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Print the current workspace name",
	Long:  `Print the current workspace name to the console.`,
	Run:   ShowRun,
	Args:  cobra.NoArgs,
}

func ShowRun(_ *cobra.Command, _ []string) {
	if viper.IsSet("settings.currentworkspace") {
		fmt.Println(viper.Get("settings.currentworkspace"))
	} else {
		panic("No current workspace! Set with 'ws switch WORKSPACE_NAME'.")
	}
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
