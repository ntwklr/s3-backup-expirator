package aws

import (
	"github.com/ntwklr/s3-backup-expirator/error"
    "github.com/aws/aws-sdk-go/aws/session"
)

func Session() (*session.Session)  {
	// Initialize a session that the SDK will use to load
    // credentials & region from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession()

	if err != nil {
		error.Exitf("Unable to create session with error: %v", err)
	}

	return sess
}