package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tech-arch1tect/DevEnv4WP/lib/boot"
	"github.com/tech-arch1tect/DevEnv4WP/lib/flags"
)

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Stops, provisions and starts the docker containers",
	Long: `Stops, provisions and starts the docker containers.
	
	Equivalent to running:
		devenv4wp stop
		devenv4wp start`,
	Run: func(cmd *cobra.Command, args []string) {
		boot.Restart()
	},
}

func init() {
	rootCmd.AddCommand(restartCmd)
	flags.RegisterCommonFlags(restartCmd)
}
