package utils

import (
	"io"
	"os"
)

func IsEmptyDir(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.ReadDir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}

func CreateDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}
