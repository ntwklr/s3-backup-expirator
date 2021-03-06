package backup

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/ntwklr/s3-backup-expirator/aws"
	"github.com/ntwklr/s3-backup-expirator/date"
	"github.com/ntwklr/s3-backup-expirator/error"
	"github.com/ntwklr/s3-backup-expirator/utilities"
	"github.com/uniplaces/carbon"
)

type Backup struct {
	ID         *string        `type:"string"`
	FilePath   *string        `type:"string"`
	Date       *carbon.Carbon `type:"timestamp"`
	ModifiedAt *time.Time     `type:"timestamp"`
	Size       *int64         `type:"integer"`
	Storage    *string        `type:"string"`
}

type Backups struct {
	Bucket  *string   `type:"string"`
	Prefix  *string   `type:"string"`
	Backups []*Backup `type:"structure"`
}

func (backups *Backups) Contains(backup *Backup) bool {
	for _, b := range backups.Backups {
		if backup == b {
			return true
		}
	}
	return false
}

type SortByDateDesc []*Backup

func (a SortByDateDesc) Len() int           { return len(a) }
func (a SortByDateDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByDateDesc) Less(i, j int) bool { return a[i].Date.GreaterThan(a[j].Date) }

type SortByFilePath []*Backup

func (a SortByFilePath) Len() int           { return len(a) }
func (a SortByFilePath) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByFilePath) Less(i, j int) bool { return *a[i].FilePath < *a[j].FilePath }

func List(bucket string, prefix *string) *Backups {
	start := time.Now()

	objectList := aws.List(bucket, *prefix).Contents

	if utilities.Bench == true {
		utilities.TimeTrack(start, "backup.List")
	}

	backups := Hydrate(objectList, prefix)

	if len(backups) < 1 {
		error.Exitf("Bucket %q is empty.", bucket)
	}

	return &Backups{Bucket: &bucket, Prefix: prefix, Backups: backups}
}

func DeleteExpired(backups *Backups, backupsToStay *Backups) {
	if utilities.Bench == true {
		defer utilities.TimeTrack(time.Now(), "backup.DeleteExpired")
	}
	for _, backup := range backups.Backups {
		if !backupsToStay.Contains(backup) {
			if !utilities.DryRun {
				aws.Delete(*backups.Bucket, *backup.FilePath)

				fmt.Printf("Object %q successfully deleted\n", *backup.FilePath)
			} else {
				fmt.Printf("Object %q will be deleted\n", *backup.FilePath)
			}
		} else {
			fmt.Printf("Object %q stays in bucket\n", *backup.FilePath)
		}
	}
}

func Hydrate(objectList []*s3.Object, prefix *string) []*Backup {
	if utilities.Bench == true {
		defer utilities.TimeTrack(time.Now(), "backup.Hydrate")
	}

	backups := []*Backup{}

	for _, item := range objectList {
		// No folders here
		if strings.HasSuffix(*item.Key, "/") {
			continue
		}
		if len(*prefix) == 0 && strings.Contains(*item.Key, "/") {
			continue
		}

		// No dot-files/folders
		if strings.HasPrefix(*item.Key, ".") {
			continue
		}
		if strings.Contains(*item.Key, "/.") {
			continue
		}

		backups = append(backups, &Backup{
			ID:         item.ETag,
			FilePath:   item.Key,
			Date:       date.Extract(item.Key),
			ModifiedAt: item.LastModified,
			Size:       item.Size,
			Storage:    item.StorageClass,
		})
	}

	if utilities.Debug == true {
		fmt.Println("backup.Hydrate:")
		utilities.PrettyPrint(backups)
	}

	return backups
}
