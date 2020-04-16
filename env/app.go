package env

import (
	"os"
)

var AppDebug = os.Getenv("APP_DEBUG")
