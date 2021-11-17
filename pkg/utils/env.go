package utils

import "os"

func IsProd() bool {
	return os.Getenv("SpecifiedConfig") == "prod"
}
