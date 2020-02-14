package cognito

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type AdminService interface {
	FindUser(string, string) (*cognitoidentityprovider.ListUsersOutput, error)
	DeleteUser(string) (*cognitoidentityprovider.AdminDeleteUserOutput, error)
}

type adminServiceImpl struct {
	provider *cognitoidentityprovider.CognitoIdentityProvider
}

func NewAdminService(provider *cognitoidentityprovider.CognitoIdentityProvider) *adminServiceImpl {
	return &adminServiceImpl{
		provider: provider,
	}
}

func (a *adminServiceImpl) FindUser(field, value string) (*cognitoidentityprovider.ListUsersOutput, error) {
	filter := fmt.Sprintf("%s=\"%s\"", field, value)
	in := &cognitoidentityprovider.ListUsersInput{
		Filter:     aws.String(filter),
		UserPoolId: aws.String(aws_user_pool_id),
	}

	return a.provider.ListUsers(in)
}

func (a *adminServiceImpl) DeleteUser(username string) (*cognitoidentityprovider.AdminDeleteUserOutput, error) {

	in := &cognitoidentityprovider.AdminDeleteUserInput{
		UserPoolId: aws.String(aws_user_pool_id),
		Username:   &username,
	}

	return a.provider.AdminDeleteUser(in)
}
