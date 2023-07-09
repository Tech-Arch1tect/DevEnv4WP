package provision

import (
	"github.com/tech-arch1tect/DevEnv4WP/lib/template"
	"github.com/tech-arch1tect/DevEnv4WP/lib/utils"
)

func ProvisionPHPConfigurations(versions []string) error {
	for _, version := range versions {
		if !utils.FileExists("./config/custom-php" + version + ".ini") {
			err := template.EmbededTemplate("custom-php.ini.tmpl", nil, "./config/custom-php"+version+".ini")
			if err != nil {
				return err
			}
		}
	}
	return nil
}
