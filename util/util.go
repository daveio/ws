package util

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"
	"strings"
)

type Shell struct {
	Name       string
	ConfigFile string
}

// create a struct
type ShellDetectionError struct {
	EnvVar string
}

// struct implements `Error` method
func (err ShellDetectionError) Error() string {
	return fmt.Sprintf("Shell detection failed. $SHELL: %s", err.EnvVar)
}

var (
	bashShell = Shell{
		Name:       "bash",
		ConfigFile: fmt.Sprintf("%s%s.bashrc", GetHomedir(), string(os.PathSeparator)),
	}
	zshShell = Shell{
		Name:       "zsh",
		ConfigFile: fmt.Sprintf("%s%s.zshrc", GetHomedir(), string(os.PathSeparator)),
	}
	fishShell = Shell{
		Name:       "fish",
		ConfigFile: fmt.Sprintf("%s%s.fishrc", GetHomedir(), string(os.PathSeparator)),
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
				fmt.Println(thisShell.ConfigFile)
				return nil, thisShell
			}
		}
	}
	return ShellDetectionError{EnvVar: currentShellName}, Shell{}
}
