package create

import (
	"errors"
	"fmt"

	"github.com/KasumiMercury/cp-cli/edit"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new Go project directory",
		Long: `Initialize a new Go project with a standardized directory structure designed for competitive programming.
Creates essential files and folders including main solution file, test cases directory, and template code.

Sets up a complete development environment ready for implementing and testing algorithmic solutions.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("create command expects 1 argument")
			}

			targetURL := args[0]

			cmd.Printf("Target URL: %s\n", targetURL)

			pID, err := ParseIDFromURL(targetURL)
			if err != nil {
				return fmt.Errorf("failed to parse target URL: %w", err)
			}

			cmd.Printf("detected Project ID: %s\n", pID)

			if err := ProjectDir(pID); err != nil {
				return fmt.Errorf("failed to create project directory: %w", err)
			}

			genPath, err := ProjectMain(pID)
			if err != nil {
				return fmt.Errorf("failed to create main.go: %w", err)
			}

			// TODO: set site
			if err := MemoryCurrentProject(pID, ""); err != nil {
				return fmt.Errorf("failed to create memory project: %w", err)
			}

			isCreateOnly, err := cmd.Flags().GetBool("create-only")
			if err != nil {
				return fmt.Errorf("failed to get 'create-only' flag: %w", err)
			}

			if isCreateOnly {
				return nil
			}

			if err := edit.OpenEditor(genPath); err != nil {
				return fmt.Errorf("failed to open editor: %w", err)
			}

			return nil
		},
	}

	createCmd.Flags().BoolP("create-only", "c", false, "only create project directory")

	return createCmd
}
