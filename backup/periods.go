package backup

import (
	"fmt"
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
	if utilities.Bench == true {
		defer utilities.TimeTrack(time.Now(), "backup.Periods")
	}

	periods := make(map[string]Period)

	startAll := start.Copy()
	endAll := startAll.SubDay().Copy()
	periods["all"] = Period{startAll, endAll}

	startDaily := endAll.Copy()
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

	if utilities.Debug == true {
		fmt.Println("backup.Periods:")
		utilities.PrettyPrint(periods)
	}

	return periods
}
