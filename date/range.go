package date

import (
	"time"
	"github.com/jinzhu/now"
)

// dateRange returns a date range function over start date to end date inclusive.
// After the end of the range, the range function returns a zero date,
// date.IsZero() is true.
func Range(start time.Time, end time.Time, interval string) func() time.Time {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	return func() time.Time {
		if start.After(end) {
			return time.Time{}
		}
		date := start

		switch interval {
		case "yearly":
			start = now.With(start.AddDate(1, 0, 0)).EndOfYear()
		case "monthly":
			start = now.With(start.AddDate(0, 1, 0)).BeginningOfMonth()
		case "weekly":
			start = now.With(start.AddDate(0, 0, 7)).EndOfWeek()
		case "daily":
			start = start.AddDate(0, 0, 1)
		}
		return date
	}
}