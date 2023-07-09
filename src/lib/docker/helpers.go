package docker

import (
	"os"

	"github.com/tech-arch1tect/DevEnv4WP/lib/configuration"
	"github.com/tech-arch1tect/DevEnv4WP/lib/utils"
)

func RunCommandContainer(command string, config configuration.Configuration) error {
	currentDir, err := os.Getwd()
	if err != nil {
		utils.DebugLog("Error getting current directory")
		return err
	}
	binds := []string{
		currentDir + "/data/:/data/",
	}
	return RunContainerWithCommand(command, config.Userid, config.Groupid, "techarchitect/devenv4wpcommandrunner:latest-"+config.Version, "devenv4wp_command", binds)

}

func RunWpCliCommand(site string, command string, config configuration.Configuration) error {
	currentDir, err := os.Getwd()
	if err != nil {
		utils.DebugLog("Error getting current directory")
		return err
	}
	binds := []string{
		currentDir + "/data/html/" + site + ":/var/www/html",
		currentDir + "/data/.wp-cli:/.wp-cli",
	}
	return RunContainerWithCommand(command, config.Userid, config.Groupid, "wordpress:cli", "devenv4wp_wpcli_"+site, binds)
}
