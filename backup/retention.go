package backup

import (
	"time"
	"github.com/jinzhu/now"
	"github.com/ntwklr/s3-backup-expirator/date"
)

func Retention(
	start time.Time, 
	retentionDaily int, 
	retentionWeekly int, 
	retentionMonthly int, 
	retentionYearly int,
	) ([]string) {
	fileRetentionDates := []string{}

	dailyStart := now.With(time.Now()).BeginningOfDay()
	dailyEnd := now.With(dailyStart.AddDate(0, 0, -(retentionDaily))).BeginningOfDay()
	
	weeklyStart := dailyEnd
	weeklyEnd := now.With(weeklyStart.AddDate(0, 0, -(7*(retentionWeekly)))).EndOfWeek()

	monthlyStart := weeklyEnd
	monthlyEnd := now.With(monthlyStart.AddDate(0, -(retentionMonthly), 0)).BeginningOfMonth()

	yearlyStart := monthlyEnd
	yearlyEnd := now.With(yearlyStart.AddDate(-(retentionYearly), 0, 0)).EndOfYear()

	for rd := date.Range(yearlyEnd, yearlyStart, "yearly"); ; {
		date := rd()
		if date.IsZero() {
			break
		}
		fileRetentionDates = append(fileRetentionDates, date.Format("2006-01-02"))
	}

	for rd := date.Range(monthlyEnd, monthlyStart, "monthly"); ; {
		date := rd()
		if date.IsZero() {
			break
		}
		fileRetentionDates = append(fileRetentionDates, date.Format("2006-01-02"))
	}

	for rd := date.Range(weeklyEnd, weeklyStart, "weekly"); ; {
		date := rd()
		if date.IsZero() {
			break
		}
		fileRetentionDates = append(fileRetentionDates, date.Format("2006-01-02"))
	}

	for rd := date.Range(dailyEnd, dailyStart, "daily"); ; {
		date := rd()
		if date.IsZero() {
			break
		}
		fileRetentionDates = append(fileRetentionDates, date.Format("2006-01-02"))
	}

	return fileRetentionDates
}