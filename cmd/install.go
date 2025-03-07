package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/daveio/ws/data"
	"github.com/daveio/ws/util"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:     "install",
	Example: `ws install`,
	Short:   "Set up your shell to work with ws",
	Long:    `Writes shell configuration for your shell of choice.`,
	Run:     InstallRun,
	Args:    cobra.NoArgs, // cobra.ExactValidArgs(1),
	// ValidArgs: []string{"bash", "zsh", "fish"},
}

func InstallRun(_ *cobra.Command, _ []string) {
	err, thisShell := util.DetectShell()
	util.Check(err)
	readBytes, err := os.ReadFile(thisShell.ConfigFile)
	util.Check(err)
	readContent := string(readBytes)
	if strings.Contains(readContent, thisShell.Template) {
		fmt.Printf("Not installing: ws env hook already installed for %s\n", thisShell.Name)
	} else {
		writeContent := []byte(fmt.Sprintf(
			"%s\n\n%s\n%s\n%s\n\n",
			readContent,
			data.HeaderComment,
			thisShell.Template,
			data.FooterComment,
		))
		backupPath := fmt.Sprintf("%s.ws.old", thisShell.ConfigFile)
		util.Check(os.WriteFile(
			backupPath, []byte(readContent), 0640))
		util.Check(os.WriteFile(
			thisShell.ConfigFile, writeContent, 0640))
		fmt.Printf("Installation complete: ws env hook installed for %s.\n", thisShell.Name)
		fmt.Println("You masffy need to restart any terminal windows, or log out and back in.")
		fmt.Printf("Your old shell configuration has been backed up to %s if you need it.\n", backupPath)
	}
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
