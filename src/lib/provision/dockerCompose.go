package provision

import (
	"os"
	"os/exec"

	"github.com/tech-arch1tect/DevEnv4WP/lib/configuration"
	"github.com/tech-arch1tect/DevEnv4WP/lib/flags"
	"github.com/tech-arch1tect/DevEnv4WP/lib/template"
)

func ProvisionDockerCompose(config configuration.Configuration) error {
	return template.EmbededTemplate("docker-compose.yml.tmpl", config, "./docker-compose.yml")
}

func StartDockerCompose() error {
	cmd := exec.Command("docker", "compose", "up", "-d")
	if flags.Debug {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
	}
	return cmd.Run()
}

func StopDockerCompose() error {
	cmd := exec.Command("docker", "compose", "down", "--remove-orphans")
	if flags.Debug {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
	}
	return cmd.Run()
}
