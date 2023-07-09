package utils

import (
	"log"

	"github.com/tech-arch1tect/DevEnv4WP/lib/flags"
)

func ExitIfError(err error) {
	if flags.Debug && err != nil {
		panic(err)
	}
	if err != nil {
		log.Println("Error encountered. Exiting.")
		log.Fatalln(err)
	}
}
