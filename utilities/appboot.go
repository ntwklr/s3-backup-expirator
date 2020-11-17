package utilities

import (
	"fmt"
	"os"
	"strconv"
)

func Boot(all int, daily int, weekly int, monthly int, yearly int) map[string]*int {
	intervals := make(map[string]*int)

	backupsAll := 2
	if len(os.Getenv("BACKUPS_ALL")) > 0 {
		value, _ := strconv.Atoi(os.Getenv("BACKUPS_ALL"))

		backupsAll = value
	}
	if all > 0 {
		backupsAll = all
	}

	intervals["all"] = &backupsAll

	backupsDaily := 8
	if len(os.Getenv("BACKUPS_DAILY")) > 0 {
		value, _ := strconv.Atoi(os.Getenv("BACKUPS_DAILY"))

		backupsDaily = value
	}
	if daily > 0 {
		backupsDaily = daily
	}

	intervals["daily"] = &backupsDaily

	backupsWeekly := 5
	if len(os.Getenv("BACKUPS_WEEKLY")) > 0 {
		value, _ := strconv.Atoi(os.Getenv("BACKUPS_WEEKLY"))

		backupsWeekly = value
	}
	if weekly > 0 {
		backupsWeekly = weekly
	}

	intervals["weekly"] = &backupsWeekly

	backupsMonthly := 13
	if len(os.Getenv("BACKUPS_MONTHLY")) > 0 {
		value, _ := strconv.Atoi(os.Getenv("BACKUPS_MONTHLY"))

		backupsMonthly = value
	}
	if monthly > 0 {
		backupsMonthly = monthly
	}
	intervals["monthly"] = &backupsMonthly

	backupsYearly := 7
	if len(os.Getenv("BACKUPS_YEARLY")) > 0 {
		value, _ := strconv.Atoi(os.Getenv("BACKUPS_YEARLY"))

		backupsYearly = value
	}
	if yearly > 0 {
		backupsYearly = yearly
	}
	intervals["yearly"] = &backupsYearly

	if Debug == true {
		fmt.Println("app.Boot:")
		PrettyPrint(intervals)
	}

	return intervals
}
