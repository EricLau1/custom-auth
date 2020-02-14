package main

import (
	"custom-auth/email"
	"custom-auth/utils"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func createCustomAuth(event *events.CognitoEventUserPoolsCreateAuthChallenge) (*events.CognitoEventUserPoolsCreateAuthChallenge, error) {
	fmt.Printf("Create Auth Challenge: %+v\n", event)

	sessions := event.Request.Session

	code := ""
	delivery := ""

	if len(sessions) > 0 {

		for _, sess := range sessions {

			challengeName := sess.ChallengeName
			challengeResult := sess.ChallengeResult
			challengeMetadata := sess.ChallengeMetadata

			fmt.Printf("[INFO] Current Session: ChallengeName: %s, ChallengeResult: %v, ChallengeMetaData: %s\n",
				challengeName, challengeResult, challengeMetadata)

			meta := strings.Split(challengeMetadata, ";")
			delivery = meta[0]
			code = meta[1]
		}
	}

	if code == "" {

		code = utils.GenerateCode()
		delivery = "email"

		user := event.Request.UserAttributes

		fmt.Printf("[INFO] User attributes: %v\n", user)

		userEmail := user["email"]

		if userEmail == "" {
			userEmail = "eric.devtt@gmail.com"
		}

		fmt.Printf("[INFO] Email: %s, Code: %s\n", userEmail, code)

		svc := email.NewEmailService()

		_, err := svc.Send(&email.Email{
			Receiver: userEmail,
			Subject:  "Código de Autenticação",
			HtmlBody: fmt.Sprintf("<p>Seu código é: <strong>%s</strong></p>", code),
			TextBody: "Autenticação Personalizada",
		})

		if err != nil {
			return event, err
		}

	}

	event.Response.PublicChallengeParameters = map[string]string{"delivery": delivery}
	event.Response.PrivateChallengeParameters = map[string]string{"answer": code}
	event.Response.ChallengeMetadata = fmt.Sprintf("%s;%s", delivery, code)

	return event, nil
}

func main() {
	lambda.Start(createCustomAuth)
}
