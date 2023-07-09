package utils

import (
	"errors"
	"regexp"
)

func GetSafeDBString(dbString string) (string, error) {
	if dbString == "" {
		return "", errors.New("empty string")
	}
	str := regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(dbString, "")
	if str == "" {
		return "", errors.New("empty string")
	}
	return str, nil
}
