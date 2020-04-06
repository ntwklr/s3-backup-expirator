package date

import (
	"fmt"
	"regexp"
)

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []string, val string) (int, bool) {
    for i, item := range slice {
		pattern := regexp.MustCompile(fmt.Sprintf(`^(.*)(%s)(.*)$`, item))

		if pattern.FindString(val) == val {
			return i, true
		}
    }
    return -1, false
}