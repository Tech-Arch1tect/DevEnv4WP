package utils

import (
	"os/user"
	"runtime"
)

func GetRunningUserIdAndGroup() (string, string, error) {
	if runtime.GOOS == "windows" {
		return "1000", "1000", nil
	}
	currentUser, err := user.Current()
	if err != nil {
		return "", "", err
	}
	return currentUser.Uid, currentUser.Gid, nil
}
