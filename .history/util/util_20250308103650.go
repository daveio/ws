package util

import (
	"fmt"
	"os"
	"strings"

	"github.com/daveio/ws/data"
	"github.com/mitchellh/go-homedir"
)

type Shell struct {
	Name       string
	ConfigFile string
	Template   string
}

type ShellDetectionError struct {
	EnvVar string
}

func (err ShellDetectionError) Error() string {
	return fmt.Sprintf("Shell detection failed. $SHELL: %s", err.EnvVar)
}

var (
	bashShell = Shell{
		Name:       "bash",
		ConfigFile: fmt.Sprintf("%s%s.bashrc", GetHomedir(), string(os.PathSeparator)),
		Template:   data.BashTemplate,
	}
	zshShell = Shell{
		Name:       "zsh",
		ConfigFile: fmt.Sprintf("%s%s.zshrc", GetHomedir(), string(os.PathSeparator)),
		Template:   data.ZshTemplate,
	}
	fishShell = Shell{
		Name: "fish",
		ConfigFile: fmt.Sprintf("%s%s.config%sfish%sconfig.fish",
			GetHomedir(),
			string(os.PathSeparator),
			string(os.PathSeparator),
			string(os.PathSeparator)),
		Template: data.FishTemplate,
	}
	shellDetectOrder = [3]Shell{fishShell, zshShell, bashShell}
)

func GetHomedir() (path string) {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	} else {
		return home
	}
}

func DetectShell() (err error, shell Shell) {
	currentShellName, ok := os.LookupEnv("SHELL")
	if ok && currentShellName != "" {
		// $SHELL is set
		for _, thisShell := range shellDetectOrder {
			if strings.Contains(currentShellName, thisShell.Name) {
				// $SHELL contains thisShell.Name
				return nil, thisShell
			}
		}
	}
	return ShellDetectionError{EnvVar: currentShellName}, Shell{}
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
