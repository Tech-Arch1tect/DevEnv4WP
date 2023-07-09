package template

import (
	"embed"
	"os"
	"text/template"
)

//go:embed templates/*
var embedos embed.FS

func EmbededTemplate(tFile string, data interface{}, output string) error {

	tmpl, err := template.New(tFile).ParseFS(embedos, "templates/*")
	if err != nil {
		return err
	}

	f, err := os.Create(output)
	if err != nil {
		return err
	}
	defer f.Close()

	err = tmpl.Execute(f, data)
	if err != nil {
		return err
	}

	return nil

}
