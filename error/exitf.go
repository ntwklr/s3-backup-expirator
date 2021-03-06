package error

import (
	"fmt"
	"os"
)

func Exitf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
