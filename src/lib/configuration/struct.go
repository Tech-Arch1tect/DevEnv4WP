package configuration

// TODO: add bind address

type Configuration struct {
	Sites   map[string]Site `yaml:"sites" validate:"required,dive,min=1"`
	Userid  int             `yaml:"userid" validate:"required"`
	Groupid int             `yaml:"groupid" validate:"required"`
	Version string          `validate:"required"`
}

type Site struct {
	Hostname    string `yaml:"hostname" validate:"required,hostname"`
	Php_version string `yaml:"php_version" validate:"required,oneof=7.4 8.0 8.1 8.2"`
}
