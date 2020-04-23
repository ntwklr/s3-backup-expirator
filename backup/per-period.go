package backup

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ntwklr/s3-backup-expirator/utilities"
)

func PerPeriod(periods map[string]*Period, backups *Backups) map[string]map[string][]*Backup {
	if utilities.Bench == true {
		defer utilities.TimeTrack(time.Now(), "backup.PerPeriod")
	}

	periodsMap := make(map[string]map[string][]*Backup)

	for periodKey, period := range periods {
		groupMap := make(map[string][]*Backup)

		for _, backup := range backups.Backups {
			if backup.Date.Between(period.Start, period.End, true) {
				groupKey := backup.Date.Format("20060102150405")

				if periodKey == "daily" {
					groupKey = backup.Date.Format("20060102")
				}

				if periodKey == "weekly" {
					year, week := backup.Date.WeekOfYear()
					groupKey = strconv.Itoa(year) + strconv.Itoa(week)
				}

				if periodKey == "monthly" {
					groupKey = backup.Date.Format("200601")
				}

				if periodKey == "yearly" {
					groupKey = backup.Date.Format("2006")
				}

				groupMap[groupKey] = append(groupMap[groupKey], backup)
			}
		}

		periodsMap[periodKey] = groupMap
	}

	if utilities.Debug == true {
		fmt.Println("backup.PerPeriod:")
		utilities.PrettyPrint(periodsMap)
	}

	return periodsMap
}
