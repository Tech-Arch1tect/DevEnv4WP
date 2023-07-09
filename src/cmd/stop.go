package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tech-arch1tect/DevEnv4WP/lib/boot"
	"github.com/tech-arch1tect/DevEnv4WP/lib/flags"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the docker containers",
	Long: `Stop the docker containers.
	
	Usage:
		devenv4wp stop`,
	Run: func(cmd *cobra.Command, args []string) {
		boot.Stop()
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
	flags.RegisterCommonFlags(stopCmd)
}
