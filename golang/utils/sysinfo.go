package utils

import (
	"os"
)

func GetSysInfo() map[string]interface{} {
	rt := map[string]interface{}{}

	rt["bin"], _ = os.Executable()
	rt["pid"] = os.Getpid()
	rt["path"], _ = os.Getwd()
	rt["host"], _ = os.Hostname()
	return rt
}
