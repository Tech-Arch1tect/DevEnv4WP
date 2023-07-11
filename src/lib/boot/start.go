package boot

import (
	"fmt"

	"github.com/tech-arch1tect/DevEnv4WP/lib/configuration"
	"github.com/tech-arch1tect/DevEnv4WP/lib/docker"
	"github.com/tech-arch1tect/DevEnv4WP/lib/provision"
	"github.com/tech-arch1tect/DevEnv4WP/lib/utils"
)

func Start() {
	conf, err := configuration.Load()
	utils.ExitIfError(err)
	fmt.Println("Configuration loaded.")
	fmt.Println("Creating Directories...")
	utils.ExitIfError(utils.CreateDir("./data/html"))
	utils.ExitIfError(utils.CreateDir("./data/certs"))
	utils.ExitIfError(utils.CreateDir("./data/.wp-cli"))
	utils.ExitIfError(utils.CreateDir("./data/mariadb"))
	utils.ExitIfError(utils.CreateDir("./config"))
	utils.ExitIfError(docker.RunCommandContainer("chmod 777 /data/.wp-cli -R", configuration.Configuration{Userid: 0, Groupid: 0, Version: conf.Version}))
	utils.ExitIfError(utils.CreateEmptyFile("./data/html/index.php"))
	fmt.Println("Provisioning PHP Configurations...")
	utils.ExitIfError(provision.ProvisionPHPConfigurations([]string{"7.4", "8.0", "8.1", "8.2"}))
	fmt.Println("Provisioning SSL Certificates...")
	utils.ExitIfError(provision.ProvisionCertificates(conf))
	fmt.Println("Provisioning Web Server Configurations...")
	for _, site := range conf.Sites {
		utils.ExitIfError(utils.CreateDir("./data/html/" + site.Hostname))
	}
	utils.ExitIfError(provision.ProvisionWeb(conf))
	fmt.Println("Provisioning Docker Compose...")
	utils.ExitIfError(provision.ProvisionDockerCompose(conf))
	fmt.Println("Provisioning Wordpress installs...")
	utils.ExitIfError(provision.ProvisionWordpress(conf))
	fmt.Println("Stopping Docker Compose Containers (if any are running)...")
	Stop()
	fmt.Println("Starting Docker Compose Containers (this may take a few minutes)...")
	utils.ExitIfError(provision.StartDockerCompose())
	fmt.Println("Provisioning wp-config.php files...")
	utils.ExitIfError(provision.ProvisionWpConfig(conf))
	fmt.Println("Provisioning Databases and Users...")
	for _, site := range conf.Sites {
		utils.ExitIfError(provision.CreateDBAndUser(conf, site.Hostname))
	}
	provision.PrintSummary(conf)
}
