package config

import (
	"github.com/spf13/viper"
	"log"
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
		log.Fatal("Unable to determine current working directory")
	}

	workspace = currentAbsolutePath + "/workspace"
}

func GetWorkspacePath() string {
	return workspace
}
