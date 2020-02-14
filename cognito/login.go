package cognito

import (
	"custom-auth/models"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type LoginService interface {
	Login(*models.User)
	CustomLogin(*models.User) (*cognitoidentityprovider.AdminInitiateAuthOutput, error)
	FinishCustomLogin(*models.User, *cognitoidentityprovider.AdminInitiateAuthOutput) (*cognitoidentityprovider.AdminRespondToAuthChallengeOutput, error)
}

type loginServiceImpl struct {
	provider *cognitoidentityprovider.CognitoIdentityProvider
}

func NewLoginService(provider *cognitoidentityprovider.CognitoIdentityProvider) *loginServiceImpl {
	return &loginServiceImpl{provider}
}

func (l *loginServiceImpl) Login(user *models.User) {

	params := map[string]*string{
		"PASSWORD": &user.Password,
		"USERNAME": &user.Nickname,
	}

	flow := "USER_PASSWORD_AUTH"

	out, err := l.initiateAuth(params, flow)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(out)
	}
}

func (l *loginServiceImpl) initiateAuth(params map[string]*string, flow string) (*cognitoidentityprovider.InitiateAuthOutput, error) {

	input := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow:       aws.String(flow),
		ClientId:       aws.String(aws_user_pool_client_id),
		AuthParameters: params,
	}

	return l.provider.InitiateAuth(input)
}

func (l *loginServiceImpl) CustomLogin(user *models.User) (*cognitoidentityprovider.AdminInitiateAuthOutput, error) {

	params := map[string]*string{
		"USERNAME": &user.Nickname,
	}

	flow := "CUSTOM_AUTH"

	input := &cognitoidentityprovider.AdminInitiateAuthInput{
		UserPoolId:     aws.String(aws_user_pool_id),
		AuthFlow:       aws.String(flow),
		ClientId:       aws.String(aws_user_pool_client_id),
		AuthParameters: params,
	}

	return l.provider.AdminInitiateAuth(input)
}

func (l *loginServiceImpl) FinishCustomLogin(user *models.User, auth *cognitoidentityprovider.AdminInitiateAuthOutput) (*cognitoidentityprovider.AdminRespondToAuthChallengeOutput, error) {
	challengeParams := make(map[string]*string)
	challengeParams["USERNAME"] = &user.Nickname
	challengeParams["ANSWER"] = &user.ConfirmationCode

	input := &cognitoidentityprovider.AdminRespondToAuthChallengeInput{
		UserPoolId:         aws.String(aws_user_pool_id),
		ChallengeName:      aws.String("CUSTOM_CHALLENGE"),
		ClientId:           aws.String(aws_user_pool_client_id),
		Session:            auth.Session,
		ChallengeResponses: challengeParams,
	}

	return l.provider.AdminRespondToAuthChallenge(input)
}
