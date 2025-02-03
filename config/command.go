package config

import "github.com/spf13/cobra"

func NewCmd() *cobra.Command {
	configCmd := &cobra.Command{
		Use:   "config",
		Short: "Manage configuration",
		Long:  "Manage configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			//TODO: implement
			return nil
		},
	}

	return configCmd
}
