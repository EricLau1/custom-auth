package email

import (
	"custom-auth/awsconf"
	"fmt"
	"log"
	"os"

	//go get -u github.com/aws/aws-sdk-go
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ses"
)

const (
	// Replace sender@example.com with your "From" address.
	// This address must be verified with Amazon SES.
	Sender = "ericlau.oliveira@gmail.com"

	// Specify a configuration set. To use a configuration
	// set, comment the next line and line 92.
	//ConfigurationSet = "ConfigSet"

	// The subject line for the email.
	// Subject = "Amazon SES Test (AWS SDK for Go)"

	// The HTML body for the email.
	// HtmlBody = "<h1>Amazon SES Test Email (AWS SDK for Go)</h1><p>This email was sent with " +
	// 	"<a href='https://aws.amazon.com/ses/'>Amazon SES</a> using the " +
	// 	"<a href='https://aws.amazon.com/sdk-for-go/'>AWS SDK for Go</a>.</p>"

	//The email body for recipients with non-HTML email clients.
	// TextBody = "This email was sent with Amazon SES using the AWS SDK for Go."

	// The character encoding for the email.
	CharSet = "UTF-8"
)

type EmailService interface {
	Send(*Email) (*ses.SendEmailOutput, error)
	Verify(string)
	ListEmails()
}

type emailServiceImpl struct {
	svc *ses.SES
}

type Email struct {
	Receiver string
	Subject  string
	HtmlBody string
	TextBody string
}

func NewEmailService() *emailServiceImpl {

	sess, err := awsconf.CreateSessionWithCredentials()
	if err != nil {
		log.Println(err)
		return nil
	}

	return &emailServiceImpl{svc: ses.New(sess)}
}

func (e *emailServiceImpl) Send(target *Email) (*ses.SendEmailOutput, error) {

	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(target.Receiver),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(target.HtmlBody),
				},
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(target.TextBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(target.Subject),
			},
		},
		Source: aws.String(Sender),
		// Uncomment to use a configuration set
		//ConfigurationSetName: aws.String(ConfigurationSet),
	}

	// Attempt to send the email.
	result, err := e.svc.SendEmail(input)

	// Display error messages if they occur.
	if err != nil {
		// if aerr, ok := err.(awserr.Error); ok {
		// 	switch aerr.Code() {
		// 	case ses.ErrCodeMessageRejected:
		// 		fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
		// 	case ses.ErrCodeMailFromDomainNotVerifiedException:
		// 		fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
		// 	case ses.ErrCodeConfigurationSetDoesNotExistException:
		// 		fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
		// 	default:
		// 		fmt.Println(aerr.Error())
		// 	}
		// } else {
		// 	// Print the error, cast err to awserr.Error to get the Code and
		// 	// Message from an error.
		// 	// fmt.Println(err.Error())
		// 	return nil, aerr.Error()
		// }

		return nil, err
	}

	// fmt.Println("Email Sent to address: " + target.Receiver)
	// fmt.Println(result)

	return result, nil
}

func (e *emailServiceImpl) Verify(receiver string) {

	// Attempt to send the email.
	out, err := e.svc.VerifyEmailAddress(&ses.VerifyEmailAddressInput{EmailAddress: aws.String(receiver)})

	// Display error messages if they occur.
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}

		return
	}

	fmt.Println(out)
}

func (e *emailServiceImpl) ListEmails() {

	result, err := e.svc.ListIdentities(&ses.ListIdentitiesInput{IdentityType: aws.String("EmailAddress")})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, e_mail := range result.Identities {
		var e_ = []*string{e_mail}

		verified, err := e.svc.GetIdentityVerificationAttributes(&ses.GetIdentityVerificationAttributesInput{Identities: e_})

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(*e_mail, verified)
		// for _, va := range verified.VerificationAttributes {
		// 	fmt.Println(va)
		// }
	}
}
