package config

import (
	"github.com/KasumiMercury/cp-cli/config/option"
	"log"
	"os"

	"github.com/spf13/viper"
)

var workspace = ""

func init() {
	if c := viper.GetString(option.WorkspaceKey); c != "" {
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
