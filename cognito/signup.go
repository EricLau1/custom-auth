package cognito

import (
	"custom-auth/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type RegisterService interface {
	SignUp(*models.User) (*cognitoidentityprovider.SignUpOutput, error)
	ConfirmSignUp(*models.User) (*cognitoidentityprovider.ConfirmSignUpOutput, error)
}

type registerServiceImpl struct {
	provider *cognitoidentityprovider.CognitoIdentityProvider
}

func NewRegisterService(provider *cognitoidentityprovider.CognitoIdentityProvider) *registerServiceImpl {
	return &registerServiceImpl{provider}
}

func (r *registerServiceImpl) SignUp(user *models.User) (*cognitoidentityprovider.SignUpOutput, error) {

	in := &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(aws_user_pool_client_id),
		Username: &user.Nickname,
		Password: &user.Password,
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			addAttr("name", user.Name),
			addAttr("nickname", user.Nickname),
			addAttr("email", user.Email),
		},
	}

	return r.provider.SignUp(in)
}

func addAttr(name, value string) *cognitoidentityprovider.AttributeType {
	return &cognitoidentityprovider.AttributeType{
		Name:  aws.String(name),
		Value: aws.String(value),
	}
}

func (r *registerServiceImpl) ConfirmSignUp(user *models.User) (*cognitoidentityprovider.ConfirmSignUpOutput, error) {

	in := &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         aws.String(aws_user_pool_client_id),
		Username:         &user.Nickname,
		ConfirmationCode: &user.ConfirmationCode,
	}

	return r.provider.ConfirmSignUp(in)
}
