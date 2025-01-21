package cmd

import (
	"github.com/KasumiMercury/cp-cli/create"
)

var createCmd = create.NewCmd()

func init() {
	rootCmd.AddCommand(createCmd)
}
