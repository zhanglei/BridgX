package utils

import (
	"path"
	"runtime"
)

func GetProjectPath() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
