package create

import (
	"errors"
	"fmt"

	"github.com/KasumiMercury/cp-cli/editor"
	"github.com/KasumiMercury/cp-cli/parse"
	"github.com/KasumiMercury/cp-cli/path"
	"github.com/spf13/cobra"
)

var ErrInvalidArg = errors.New("invalid argument")

func NewCmd() *cobra.Command {
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new Go project directory",
		Long: `Initialize a new Go project with a standardized directory structure designed for competitive programming.
Creates essential files and folders including main solution file, test cases directory, and template code.

Sets up a complete development environment ready for implementing and testing algorithmic solutions.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return ErrInvalidArg
			}

			targetURL := args[0]

			cmd.Printf("Target URL: %s\n", targetURL)

			err := createCmd(cmd, targetURL)
			if err != nil {
				return err
			}

			return nil
		},
	}

	createCmd.Flags().BoolP("create-only", "c", false, "only create project directory")

	return createCmd
}

func createCmd(cmd *cobra.Command, targetURL string) error {
	pID, _, err := parse.IDFromURL(targetURL)
	if err != nil {
		return fmt.Errorf("failed to parse target URL: %w", err)
	}

	cmd.Printf("detected Project ID: %s\n", pID)

	// TODO: use config
	projectRoot := "./project"

	projectPath, err := path.ProjectDirPath(projectRoot, pID)
	if err != nil {
		return fmt.Errorf("failed to get project dir: %w", err)
	}

	if err := ProjectDir(projectPath, pID); err != nil {
		return err
	}

	// TODO: set site
	if err := MemoryCurrentProject(pID, ""); err != nil {
		return fmt.Errorf("failed memory current problem: %w", err)
	}

	isCreateOnly, err := cmd.Flags().GetBool("create-only")
	if err != nil {
		return fmt.Errorf("failed flag get create-only value: %w", err)
	}

	if isCreateOnly {
		return nil
	}

	if err := editor.Open(
		"", projectPath, "main.go",
		cmd.InOrStdin(), cmd.OutOrStdout(), cmd.OutOrStderr(),
	); err != nil {
		return fmt.Errorf("failed to open editor: %w", err)
	}

	return nil
}
