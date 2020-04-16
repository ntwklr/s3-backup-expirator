package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/ntwklr/s3-backup-expirator/error"
)

// List get the list of objects from s3
func List(bucket string, prefix string) *s3.ListObjectsV2Output {
	session := Session()
	client := Client(session)

	input := &s3.ListObjectsV2Input{
		Bucket:    aws.String(bucket),
		Prefix:    aws.String(prefix),
		Delimiter: aws.String("/"),
	}

	// Get the list of items
	response, err := client.ListObjectsV2(input)
	if err != nil {
		error.Exitf("Unable to list items in bucket %q, %v", bucket, err)
	}

	return response
}
