package utils

import (
	"log"

	"github.com/tech-arch1tect/DevEnv4WP/lib/flags"
)

func DebugLog(msg string) {
	if flags.Debug {
		log.Println(msg)
	}
}
