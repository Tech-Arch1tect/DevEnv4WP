package configuration

type Configuration struct {
	Sites       map[string]Site `yaml:"sites" validate:"required,min=1,dive"`
	Userid      int             `yaml:"userid" validate:"required"`
	Groupid     int             `yaml:"groupid" validate:"required"`
	Version     string          `validate:"required"`
	BindAddress string          `yaml:"bind_address" validate:"required,ipv4"`
	WebServer   string          `yaml:"web_server" validate:"required,oneof=nginx apache"`
}

type Site struct {
	Hostname    string `yaml:"hostname" validate:"required,hostname"`
	Php_version string `yaml:"php_version" validate:"required,oneof=7.4 8.0 8.1 8.2 8.3 8.4"`
}
