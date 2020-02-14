package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func verifyCustomAuth(event *events.CognitoEventUserPoolsVerifyAuthChallenge) (*events.CognitoEventUserPoolsVerifyAuthChallenge, error) {
	fmt.Printf("Verify Auth Challenge: %+v\n", event)

	params := event.Request.PrivateChallengeParameters
	userParams := event.Request.UserAttributes

	fmt.Printf("[INFO] User Params: %v\n", userParams)

	expectedAnswer := params["answer"]
	answer, ok := event.Request.ChallengeAnswer.(string)
	if ok {
		event.Response.AnswerCorrect = answer == expectedAnswer
	}

	return event, nil
}

func main() {
	lambda.Start(verifyCustomAuth)
}
