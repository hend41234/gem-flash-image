package utils

import (
	"os"
)

func NotExistPath(path string) bool {
	// workDir, _ := os.Getwd()
	_, err := os.Stat(path)
	return !os.IsExist(err)
}

func CreatePath(path string) {
	if !NotExistPath(path) {
		return
	}
	os.Mkdir(path, 0755)
}
