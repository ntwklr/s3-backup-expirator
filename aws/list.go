package aws

import (
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/s3"
	"github.com/etenzy/s3-backup-expirator/error"
)

func List(bucket string) (*s3.ListObjectsV2Output) {
	session := Session();
	client := Client(session)

	// Get the list of items
	response, err := client.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucket)})
    if err != nil {
        error.Exitf("Unable to list items in bucket %q, %v", bucket, err)
	}

	return response
}