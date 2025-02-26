package config

import (
	"github.com/spf13/viper"
	"os"
)

var workspace = ""

func init() {
	if c := viper.GetString("workspace"); c != "" {
		workspace = c

		return
	}

	currentAbsolutePath, err := os.Getwd()
	if err != nil {
		// TODO: improve error handling
		panic(err)
	}

	workspace = currentAbsolutePath + "/workspace"
}

func GetWorkspacePath() string {
	return workspace
}
