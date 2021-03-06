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
func Periods(start *carbon.Carbon, intervals map[string]*int) map[string]*Period {
	if utilities.Bench == true {
		defer utilities.TimeTrack(time.Now(), "backup.Periods")
	}

	periods := make(map[string]*Period)

	startAll := start.EndOfDay()
	endAll := start.SubDays(*intervals["all"])
	periods["all"] = &Period{startAll, endAll}

	startDaily := endAll.EndOfDay()
	endDaily := startDaily.SubDays(*intervals["daily"]).StartOfDay()
	periods["daily"] = &Period{startDaily, endDaily}

	startWeekly := endDaily.EndOfWeek()
	endWeekly := startWeekly.SubWeeks(*intervals["weekly"]).StartOfWeek()
	periods["weekly"] = &Period{startWeekly, endWeekly}

	startMonthly := endWeekly.EndOfMonth()
	endMonthly := startMonthly.SubMonths(*intervals["monthly"]).StartOfMonth()
	periods["monthly"] = &Period{startMonthly, endMonthly}

	startYearly := endMonthly.EndOfYear()
	endYearly := startYearly.SubYears(*intervals["yearly"]).StartOfYear()
	periods["yearly"] = &Period{startYearly, endYearly}

	if utilities.Debug == true {
		fmt.Println("backup.Periods:")
		utilities.PrettyPrint(periods)
	}

	return periods
}
