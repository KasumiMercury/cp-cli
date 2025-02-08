package cmd

import (
	"github.com/KasumiMercury/cp-cli/config"
)

var configCmd = config.NewCmd()

func init() {
	rootCmd.AddCommand(configCmd)
}
