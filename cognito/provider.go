package cognito

import (
	"custom-auth/awsconf"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

var (
	aws_user_pool_id        = ""
	aws_user_pool_client_id = ""
)

func New() *cognitoidentityprovider.CognitoIdentityProvider {

	aws_user_pool_id = os.Getenv("AWS_USER_POOL_ID")
	aws_user_pool_client_id = os.Getenv("AWS_POOL_CLIENT_ID")

	sess, err := awsconf.CreateSessionWithCredentials()
	if err != nil {
		log.Fatal(err)
	}

	return cognitoidentityprovider.New(sess)
}
