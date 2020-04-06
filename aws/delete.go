package aws

import (
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/s3"
	"github.com/etenzy/s3-backup-expirator/error"
)

func Delete(bucket string, object string) (*s3.DeleteObjectOutput) {
	session := Session();
	client := Client(session)

	// Delete the item
	response, err := client.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(bucket), Key: aws.String(object)})
	if err != nil {
		error.Exitf("Unable to delete object %q from bucket %q, %v", object, bucket, err)
	}

	err = client.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(object),
	})
	if err != nil {
		error.Exitf("Error occurred while waiting for object %q to be deleted, %v", object, err)
	}

	return response
}