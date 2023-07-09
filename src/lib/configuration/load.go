package configuration

import (
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/tech-arch1tect/DevEnv4WP/lib/flags"
	"gopkg.in/yaml.v3"
)

func Load() (Configuration, error) {

	file, err := os.ReadFile("devenv4wp.yaml")
	if err != nil {
		return Configuration{}, err
	}

	configuration := Configuration{}

	err = yaml.Unmarshal(file, &configuration)
	if err != nil {
		return Configuration{}, err
	}

	configuration.Version = Version()

	err = configuration.Validate()

	return configuration, err
}

func (c *Configuration) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

func Version() string {
	if flags.Dev {
		return "dev"
	}
	return "v0.0.1"
}
