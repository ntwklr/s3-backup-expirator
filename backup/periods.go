package backup

import (
	"time"

	"github.com/ntwklr/s3-backup-expirator/utilities"
	"github.com/uniplaces/carbon"
)

// Period of datetimes
type Period struct {
	Start *carbon.Carbon
	End   *carbon.Carbon
}

// Periods calculates the retention periods
func Periods(start carbon.Carbon, intervals map[string]int) map[string]Period {
	end := time.Now()

	periods := make(map[string]Period)

	startDaily := start.Copy()
	endDaily := startDaily.SubDays(intervals["daily"]).Copy()
	periods["daily"] = Period{startDaily, endDaily}

	startWeekly := endDaily.Copy()
	endWeekly := startWeekly.SubWeeks(intervals["weekly"]).Copy()
	periods["weekly"] = Period{startWeekly, endWeekly}

	startMonthly := endWeekly.Copy()
	endMonthly := startMonthly.SubMonths(intervals["monthly"]).Copy()
	periods["monthly"] = Period{startMonthly, endMonthly}

	startYearly := endMonthly.Copy()
	endYearly := startYearly.SubYears(intervals["yearly"]).Copy()
	periods["yearly"] = Period{startYearly, endYearly}

	if utilities.Explain == true {
		utilities.TimeTrack(end, "backup.Periods")
	}

	if utilities.Explain == true {
		utilities.PrettyPrint(periods)
	}

	return periods
}
