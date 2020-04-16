package backup

import (
	"sort"
	"time"

	"github.com/ntwklr/s3-backup-expirator/utilities"
)

type SortByValueDesc []string

func (a SortByValueDesc) Len() int           { return len(a) }
func (a SortByValueDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByValueDesc) Less(i, j int) bool { return a[i] > a[j] }

func RemoveForAllPeriodsExceptOne(backupsPerPeriod map[string]map[string][]Backup, intervals map[string]int) Backups {
	if utilities.Explain == true {
		defer utilities.TimeTrack(time.Now(), "backup.RemoveForAllPeriodsExceptOne")
	}

	backups := []Backup{}

	for periodKey, period := range backupsPerPeriod {
		groupKeys := make([]string, 0, len(period))
		for k := range period {
			groupKeys = append(groupKeys, k)
		}

		sort.Sort(SortByValueDesc(groupKeys))

		interval := intervals[periodKey]
		groupKeysLength := len(groupKeys)

		if groupKeysLength < interval {
			interval = groupKeysLength
		}

		groupKeys = groupKeys[:interval]

		for _, groupKey := range groupKeys {
			group := period[groupKey]

			sort.Sort(SortByDateDesc(group))

			group = group[:1]

			for _, backup := range group {
				backups = append(backups, backup)
			}
		}
	}

	sort.Sort(SortByFilePath(backups))

	return Backups{Backups: backups}
}
