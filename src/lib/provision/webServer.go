package provision

import (
	"github.com/tech-arch1tect/DevEnv4WP/lib/configuration"
	"github.com/tech-arch1tect/DevEnv4WP/lib/docker"
	"github.com/tech-arch1tect/DevEnv4WP/lib/template"
	"github.com/tech-arch1tect/DevEnv4WP/lib/utils"
)

func ProvisionWeb(conf configuration.Configuration) error {
	if conf.WebServer == "nginx" {
		return ProvisionNginx(conf.Sites)
	} else if conf.WebServer == "apache" {
		return ProvisionApache(conf.Sites)
	}
	return nil
}

func ProvisionApache(sites map[string]configuration.Site) error {
	return template.EmbededTemplate("apache.conf.tmpl", sites, "data/apache-hostnames.conf")
}

func ProvisionNginx(sites map[string]configuration.Site) error {
	return template.EmbededTemplate("nginx.conf.tmpl", sites, "data/nginx-hostnames.conf")
}

func GenerateSelfSignedCertificate(hostname string, conf configuration.Configuration) error {
	return docker.RunCommandContainer("openssl req -x509 -nodes -days 3650 -newkey rsa:2048 -keyout ./data/certs/"+hostname+".key -out ./data/certs/"+hostname+".crt -subj /C=US/ST=DevEnv4WP/L=DevEnv4WP/O=DevEnv4WP./CN="+hostname, conf)
}

func ProvisionCertificates(conf configuration.Configuration) error {
	for _, site := range conf.Sites {
		if utils.FileExists("data/certs/"+site.Hostname+".crt") && utils.FileExists("data/certs/"+site.Hostname+".key") {
			continue
		}
		err := GenerateSelfSignedCertificate(site.Hostname, conf)
		if err != nil {
			return err
		}
	}
	// phpmyadmin
	if !utils.FileExists("data/certs/pma.local.crt") || !utils.FileExists("data/certs/pma.local.key") {
		err := GenerateSelfSignedCertificate("pma.local", conf)
		if err != nil {
			return err
		}
	}
	// mailpit
	if !utils.FileExists("data/certs/mailpit.local.crt") || !utils.FileExists("data/certs/mailpit.local.key") {
		err := GenerateSelfSignedCertificate("mailpit.local", conf)
		if err != nil {
			return err
		}
	}

	return nil
}
