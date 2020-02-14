package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func defineCustomAuth(event *events.CognitoEventUserPoolsDefineAuthChallenge) (*events.CognitoEventUserPoolsDefineAuthChallenge, error) {
	fmt.Printf("Define Auth Challenge: %+v\n", event)

	sessions := event.Request.Session
	failed := len(sessions) > 3

	if len(sessions) > 0 && !failed {

		for _, sess := range sessions {

			fmt.Printf("[INFO] Challenge Metadata: %v\n", sess.ChallengeMetadata)
			fmt.Printf("[INFO] Challenge Name: %v\n", sess.ChallengeName)
			fmt.Printf("[INFO] Challenge Result: %v\n", sess.ChallengeResult)

			if sess.ChallengeResult {
				event.Response.ChallengeName = "CUSTOM_CHALLENGE"
				event.Response.IssueTokens = true
				event.Response.FailAuthentication = false

				return event, nil
			}
		}
	}

	event.Response.ChallengeName = "CUSTOM_CHALLENGE"
	event.Response.IssueTokens = false
	event.Response.FailAuthentication = failed

	return event, nil
}

func main() {
	lambda.Start(defineCustomAuth)
}
