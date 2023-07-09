package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "DevEnv4WP",
	Short: "Manage your WordPress development environment",
	Long:  `DevEnv4WP is a CLI tool to manage your WordPress development environment.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
