package provision

import (
	"fmt"

	"github.com/tech-arch1tect/DevEnv4WP/lib/configuration"
)

func PrintSummary(conf configuration.Configuration) {
	sitesString := ""
	for _, site := range conf.Sites {
		sitesString += " " + site.Hostname
	}
	fmt.Println("\nServices are provisioned and listening on " + conf.BindAddress + ".")
	fmt.Println("IMPORTANT: Your hosts file should have the following entry:")
	fmt.Println(conf.BindAddress + " devenv4wp.local" + sitesString)

	fmt.Println("\nOnce you have added the above entry to your hosts file, you can access the following services:")
	fmt.Println("phpMyAdmin: https://devenv4wp.local/phpmyadmin/")
	fmt.Println("Mailpit: https://devenv4wp.local/mailpit/")
	for _, site := range conf.Sites {
		fmt.Println(site.Hostname + ": https://" + site.Hostname)
	}

}
