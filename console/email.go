package console

import (
	"custom-auth/email"
	"fmt"
	"log"
)

const (
	SEND_EMAIL   = "e"
	VERIFY_EMAIL = "v"
	LIST_EMAILS  = "l"
)

func (c *consoleImpl) sendEmail() {
	e := &email.Email{}
	fmt.Print("Enviar Para: ")
	fmt.Scan(&e.Receiver)
	fmt.Print("Assunto: ")
	fmt.Scan(&e.Subject)
	fmt.Print("Mensagem: ")
	fmt.Scan(&e.TextBody)

	e.HtmlBody = "<h1>" + e.TextBody + "!</h1>"

	out, err := c.emails.Send(e)
	if err != nil {
		log.Println(err)
	} else {
		JSON(out)
	}
}

func (c *consoleImpl) verifyEmail() {
	var e string
	fmt.Print("Email? ")
	fmt.Scan(&e)

	c.emails.Verify(e)
}

func (c *consoleImpl) listEmails() {
	c.emails.ListEmails()
}
