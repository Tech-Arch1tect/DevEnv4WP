package boot

import (
	"fmt"

	"github.com/tech-arch1tect/DevEnv4WP/lib/provision"
	"github.com/tech-arch1tect/DevEnv4WP/lib/utils"
)

func Stop() {
	fmt.Println("Stopping the docker containers...")
	utils.ExitIfError(provision.StopDockerCompose())
	fmt.Println("Docker containers stopped.")
}
