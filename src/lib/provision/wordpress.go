package provision

import (
	"fmt"

	"github.com/tech-arch1tect/DevEnv4WP/lib/configuration"
	"github.com/tech-arch1tect/DevEnv4WP/lib/docker"
	"github.com/tech-arch1tect/DevEnv4WP/lib/utils"
)

func ProvisionWordpress(conf configuration.Configuration) error {
	for _, site := range conf.Sites {
		isEmpty, err := utils.IsEmptyDir("data/html/" + site.Hostname)
		if err != nil {
			return err
		}
		if isEmpty {
			fmt.Println(site.Hostname + " has no WP install. Using WP-CLI to download WP core files to data/html/" + site.Hostname)
			err := docker.RunWpCliCommand(site.Hostname, "core download", conf)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func ProvisionWpConfig(conf configuration.Configuration) error {
	for _, site := range conf.Sites {
		if utils.FileExists("data/html/" + site.Hostname + "/wp-config.php") {
			utils.DebugLog("wp-config.php already exists for " + site.Hostname + ". Skipping wp-config.php generation")
			continue
		}
		sDb, err := utils.GetSafeDBString(site.Hostname)
		if err != nil {
			return err
		}
		err = docker.RunWpCliCommand(site.Hostname, "config create --dbname="+sDb+" --dbuser="+sDb+" --dbpass="+sDb+" --dbhost=mariadb --dbprefix=wp_ --skip-check", conf)
		if err != nil {
			return err
		}
	}
	return nil
}
