package set

import "github.com/spf13/cobra"

func NewCmd() *cobra.Command {
	setCmd := &cobra.Command{
		Use:   "set <key> <value>",
		Short: "Set configuration",
		Long:  "Set configuration",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			set(cmd, args)
		},
	}

	return setCmd
}

func set(cmd *cobra.Command, args []string) {}
