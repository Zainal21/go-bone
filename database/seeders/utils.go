package seeders

import (
	"path"
	"runtime"
)

func GetSourcePath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
