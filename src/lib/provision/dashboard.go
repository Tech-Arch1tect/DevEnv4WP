package provision

import (
	"github.com/tech-arch1tect/DevEnv4WP/lib/configuration"
	"github.com/tech-arch1tect/DevEnv4WP/lib/template"
)

func ProvisionDashboard(conf configuration.Configuration) error {
	return template.EmbededTemplate("dashboard.html.tmpl", conf.Sites, "./data/dashboard/index.html")
}
