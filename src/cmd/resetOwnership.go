package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tech-arch1tect/DevEnv4WP/lib/configuration"
	"github.com/tech-arch1tect/DevEnv4WP/lib/docker"
	"github.com/tech-arch1tect/DevEnv4WP/lib/flags"
	"github.com/tech-arch1tect/DevEnv4WP/lib/utils"
)

var resetOwnershipCmd = &cobra.Command{
	Use:   "resetOwnership",
	Short: "Reset ownership of files and directories",
	Long:  `Reset ownership of files and directories to the user and group defined in the configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Resetting ownership of files and directories...")
		conf, err := configuration.Load()
		utils.ExitIfError(err)
		utils.ExitIfError(docker.RunCommandContainer("chown "+strconv.Itoa(conf.Userid)+":"+strconv.Itoa(conf.Groupid)+" /data/ -R", configuration.Configuration{Userid: 0, Groupid: 0, Version: conf.Version}))
		fmt.Println("Ownership reset.")
	},
}

func init() {
	rootCmd.AddCommand(resetOwnershipCmd)
	flags.RegisterCommonFlags(resetOwnershipCmd)
}
