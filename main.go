package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/uniplaces/carbon"

	"github.com/ntwklr/s3-backup-expirator/backup"
	"github.com/ntwklr/s3-backup-expirator/date"
	"github.com/ntwklr/s3-backup-expirator/error"
	"github.com/ntwklr/s3-backup-expirator/utilities"
)

const (
	projectOwner   = "ntwklr"
	projectRepo    = "s3-backup-expirator"
	binaryPlatform = runtime.GOOS + "_" + runtime.GOARCH
)

var version string

func init() {
	dotenv := ".env"
	appenv := os.Getenv("APP_ENV")

	if appenv != "" {
		dotenv = dotenv + "." + strings.ToLower(appenv)
	}

	errLoad := godotenv.Overload("./" + dotenv)
	if errLoad != nil {
		_, errStat := os.Stat("./" + dotenv)

		if os.IsNotExist(errStat) {
			example, errRead := godotenv.Read("./.env.example")

			if errRead != nil {
				log.Fatal(errRead)
			}

			errWrite := godotenv.Write(example, "./"+dotenv)

			if errWrite != nil {
				log.Fatal(errWrite)
			}
		} else {
			log.Fatal(errLoad)
		}
	}
}

// Deletes the specified object in the specified S3 Bucket in the region configured in the shared config
// or AWS_REGION environment variable.
//
// Usage:
//    go run s3-backup-expirator BUCKET_NAME
func main() {
	bootStart := time.Now()

	if version == "" {
		version = "0.0.0"
	}

	if len(os.Args) < 2 {
		error.Exitf("Bucket name required\nUsage: %s bucket_name",
			os.Args[0])
	}

	if len(os.Args) > 1 && os.Args[1] == "self-update" {
		utilities.Update(projectOwner, projectRepo, binaryPlatform, version)
	}

	if len(os.Args) > 1 && (os.Args[1] == "-V" || os.Args[1] == "--version") {
		fmt.Println("Version: " + version)
		os.Exit(0)
	}

	bucket := os.Args[len(os.Args)-1]

	prefix := flag.String("prefix", "", "File-Prefix")
	startDate := flag.String("start-date", "", "Start-Date for Backups (default: now)")
	daily := flag.Int("daily", 0, "Daily Backup Retention Policy.")
	weekly := flag.Int("weekly", 0, "Weekly Backup Retention Policy.")
	monthly := flag.Int("monthly", 0, "Monthly Backup Retention Policy.")
	yearly := flag.Int("yearly", 0, "Yearly Backup Retention Policy.")
	bench := flag.Bool("bench", false, "Print the backup calculations.")
	debug := flag.Bool("debug", false, "Print the results of the backup calculations.")
	dryRun := flag.Bool("dry-run", false, "Print the commands that would be executed, but do not execute them.")
	flag.Parse()

	utilities.Bench = *bench
	utilities.Debug = *debug
	utilities.DryRun = *dryRun

	backupsStart := carbon.Now()

	if *startDate != "" {
		backupsStart = date.Extract(startDate)
	}

	periodIntervals := utilities.Boot(*daily, *weekly, *monthly, *yearly)

	if utilities.Bench == true {
		utilities.TimeTrack(bootStart, "app.Boot")
	}
	appStart := time.Now()

	backups := backup.List(bucket, prefix)

	periods := backup.Periods(backupsStart, periodIntervals)

	backupsPerPeriod := backup.PerPeriod(periods, backups)

	backupsToStay := backup.RemoveForAllPeriodsExceptOne(backupsPerPeriod, periodIntervals)

	backup.DeleteExpired(backups, backupsToStay)

	if utilities.Bench == true {
		utilities.TimeTrack(appStart, "app.Execute")
	}
}
