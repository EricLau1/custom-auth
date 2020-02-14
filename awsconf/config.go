package awsconf

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func CreateSessionWithCredentials() (*session.Session, error) {

	aws_access_key := os.Getenv("CONSOLE_AWS_ACCESS_KEY")
	aws_secret_key := os.Getenv("CONSOLE_AWS_SECRET_KEY")

	creds := credentials.NewStaticCredentials(aws_access_key, aws_secret_key, "")

	return session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("CONSOLE_AWS_REGION")),
		Credentials: creds,
	})
}
