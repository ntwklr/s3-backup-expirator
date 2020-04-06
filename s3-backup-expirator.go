package main

import (
	"os"
	"fmt"
	"flag"
	"time"
	"github.com/jinzhu/now"
	"github.com/etenzy/s3-backup-expirator/aws"
	"github.com/etenzy/s3-backup-expirator/backup"
	"github.com/etenzy/s3-backup-expirator/date"
	"github.com/etenzy/s3-backup-expirator/error"
)

// Deletes the specified object in the specified S3 Bucket in the region configured in the shared config
// or AWS_REGION environment variable.
//
// Usage:
//    go run s3-backup-expirator BUCKET_NAME
func main()  {
	if len(os.Args) < 2 {
        error.Exitf("Bucket name required\nUsage: %s bucket_name",
            os.Args[0])
	}

	now.WeekStartDay = time.Monday

	retentionDaily := flag.Int("daily", 8, "Daily Backup Retention Policy.")
	retentionWeekly := flag.Int("weekly", 5, "Weekly Backup Retention Policy.")
	retentionMonthly := flag.Int("monthly", 13, "Monthly Backup Retention Policy.")
	retentionYearly := flag.Int("yearly", 7, "Yearly Backup Retention Policy.")
	explainMode := flag.Bool("explain", false, "Explains wich files retain in bucket.")
	dryRun := flag.Bool("dry-run", false, "Print the commands that would be executed, but do not execute them.")
	flag.Parse()
	
	bucket := os.Args[len(os.Args)-1]

	fileRetentionDates := backup.Retention(
		now.BeginningOfDay(), 
		*retentionDaily, 
		*retentionWeekly, 
		*retentionMonthly, 
		*retentionYearly,
	)

	if *explainMode == true {
		for i, item := range fileRetentionDates {
			if i >= 0 {
				fmt.Println(item)
			}
		}
	}

	for _, item := range aws.List(bucket).Contents {
		k, found := date.Find(fileRetentionDates, *item.Key)
		if !found || k < 0 {
			if ! *dryRun {
				aws.Delete(bucket, *item.Key)

				fmt.Printf("Object %q successfully deleted\n", *item.Key)
			} else {
				fmt.Printf("Object %q will be deleted\n", *item.Key)
			}
		} else {
			fmt.Printf("Object %q stays in bucket\n", *item.Key)
		}
    }
}