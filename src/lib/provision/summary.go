package provision

import (
	"fmt"

	"github.com/tech-arch1tect/DevEnv4WP/lib/configuration"
)

func PrintSummary(sites map[string]configuration.Site) {
	sitesString := ""
	for _, site := range sites {
		sitesString += " " + site.Hostname
	}
	fmt.Println("\nServices are provisioned and listening on localhost.")
	fmt.Println("IMPORTANT: Your hosts file should have the following entry:")
	fmt.Println("127.0.0.1 pma.local mailpit.local" + sitesString)

	fmt.Println("\nOnce you have added the above entry to your hosts file, you can access the following services:")
	fmt.Println("phpMyAdmin: https://pma.local")
	fmt.Println("Mailpit: https://mailpit.local")
	for _, site := range sites {
		fmt.Println(site.Hostname + ": https://" + site.Hostname)
	}

}
