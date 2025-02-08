package config

import (
	"github.com/KasumiMercury/cp-cli/config/list"
	"github.com/KasumiMercury/cp-cli/config/set"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	configCmd := &cobra.Command{
		Use:   "config",
		Short: "Manage configuration",
		Long:  "Manage configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	configCmd.AddCommand(list.NewCmd())
	configCmd.AddCommand(set.NewCmd())

	return configCmd
}
