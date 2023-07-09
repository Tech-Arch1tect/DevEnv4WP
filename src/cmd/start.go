package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tech-arch1tect/DevEnv4WP/lib/boot"
	"github.com/tech-arch1tect/DevEnv4WP/lib/flags"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Provision and start the docker containers",
	Long: `Provision and start the docker containers. 
	Uses the configuration file (devenv4wp.yaml) in the current directory.
	
	Usage:
		devenv4wp start`,
	Run: func(cmd *cobra.Command, args []string) {
		boot.Start()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	flags.RegisterCommonFlags(startCmd)
}
