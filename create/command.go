package create

import (
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new Go project directory",
		Long: `Initialize a new Go project with a standardized directory structure designed for competitive programming.
Creates essential files and folders including main solution file, test cases directory, and template code.

Sets up a complete development environment ready for implementing and testing algorithmic solutions.`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			// TODO: implement
			return nil
		},
	}

	createCmd.Flags().StringP("target", "t", "", "target URL")
	if err := createCmd.MarkFlagRequired("target"); err != nil {
		panic("failed to mark flag 'target'" + err.Error())
	}

	return createCmd
}
