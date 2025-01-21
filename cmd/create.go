/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/KasumiMercury/cp-cli/create"
)

var createCmd = create.NewCmd()

// createCmd represents the create command.
//var createCmd = &cobra.Command{
//	Use:   "create",
//	Short: "Create a new Go project directory",
//	Long: `Initialize a new Go project with a standardized directory structure designed for competitive programming.
//Creates essential files and folders including main solution file, test cases directory, and template code.
//
//Sets up a complete development environment ready for implementing and testing algorithmic solutions.`,
//	Run: func(_ *cobra.Command, _ []string) {
//		pID, err := create.ParseIDFromURL(targetURL)
//		if err != nil {
//			slog.Error(err.Error())
//		}
//
//		fmt.Println(pID)
//
//		if err := create.ProjectDir(pID); err != nil {
//			slog.Error(err.Error())
//		}
//
//		genPath, err := create.ProjectMain(pID)
//		if err != nil {
//			slog.Error(err.Error())
//		}
//
//		if err := edit.OpenEditor(genPath); err != nil {
//			slog.Error(err.Error())
//		}
//	},
//}

func init() {
	rootCmd.AddCommand(createCmd)
}
