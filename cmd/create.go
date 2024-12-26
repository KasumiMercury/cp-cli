/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/KasumiMercury/cp-cli/create"
	"github.com/spf13/cobra"
	"log/slog"
)

var targetUrl string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Go project directory",
	Long: `Initialize a new Go project with a standardized directory structure designed for competitive programming.
Creates essential files and folders including main solution file, test cases directory, and template code.

Sets up a complete development environment ready for implementing and testing algorithmic solutions.`,
	Run: func(cmd *cobra.Command, args []string) {
		pId, err := create.ParseIdFromUrl(targetUrl)
		if err != nil {
			slog.Error(err.Error())
		}

		fmt.Println(pId)

		if err := create.ProjectDir(pId); err != nil {
			slog.Error(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createCmd.Flags().StringVarP(&targetUrl, "target", "t", "", "Target URL")
}
