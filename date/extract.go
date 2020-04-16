package date

import (
	"fmt"
	"regexp"

	"github.com/uniplaces/carbon"
)

// Extract extracts the datetime from string
func Extract(val string) *carbon.Carbon {
	format := "2006-01-02-15-04-05"

	pattern := regexp.MustCompile(`(\d{4}-\d{2}-\d{2})(-\d{2})?(-\d{2})?(-\d{2})?`)
	timeString := pattern.FindString(val)

	if len(timeString) == 16 {
		format = "2006-01-02-15-04"
	}

	if len(timeString) == 13 {
		format = "2006-01-02-15"
	}

	if len(timeString) == 10 {
		format = "2006-01-02"
	}

	date, err := carbon.CreateFromFormat(format, timeString, carbon.Now().TimeZone())
	if err != nil {
		fmt.Println("Error while parsing date :", err)
	}

	return date
}
