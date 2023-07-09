package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tech-arch1tect/DevEnv4WP/lib/flags"
	"github.com/tech-arch1tect/DevEnv4WP/lib/template"
	"github.com/tech-arch1tect/DevEnv4WP/lib/utils"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the configuration file for DevEnv4WP",
	Long: `Initialize the configuration file for DevEnv4WP.
	
	Usage:
		devenv4wp init
		# This will create a configuration file in the current directory.
		# The configuration file will be named devenv4wp.yaml.`,
	Run: func(cmd *cobra.Command, args []string) {
		userID, groupID, err := utils.GetRunningUserIdAndGroup()
		utils.ExitIfError(err)
		data := map[string]interface{}{
			"Userid":  userID,
			"Groupid": groupID,
		}
		utils.ExitIfError(template.EmbededTemplate("devenv4wp.yaml.tmpl", data, "devenv4wp.yaml"))
		fmt.Println("devenv4wp.yaml created. Init complete.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	flags.RegisterCommonFlags(initCmd)
}
