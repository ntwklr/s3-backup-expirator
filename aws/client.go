package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func Client(session *session.Session) (*s3.S3) {
	// Create S3 service client
	return s3.New(session)
}