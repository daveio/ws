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
	"github.com/daveio/ws/data"
	"github.com/daveio/ws/util"
	"io/ioutil"
	"strings"

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

func InstallRun(cmd *cobra.Command, args []string) {
	err, thisShell := util.DetectShell()
	util.Check(err)
	readBytes, err := ioutil.ReadFile(thisShell.ConfigFile)
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
		util.Check(ioutil.WriteFile(
			backupPath, []byte(readContent), 0640))
		util.Check(ioutil.WriteFile(
			thisShell.ConfigFile, writeContent, 0640))
		fmt.Printf("Installation complete: ws env hook installed for %s.\n", thisShell.Name)
		fmt.Println("You may need to restart any terminal windows, or log out and back in.")
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
