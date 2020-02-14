package console

import (
	"custom-auth/models"
	"fmt"
	"log"
)

const (
	CREATE_USER  = "1"
	CONFIRM_USER = "2"
)

func CreateUser() *models.User {
	user := &models.User{}
	fmt.Print("Name? ")
	fmt.Scan(&user.Name)
	fmt.Print("Email? ")
	fmt.Scan(&user.Email)
	fmt.Print("Username? ")
	fmt.Scan(&user.Nickname)
	fmt.Print("Password? ")
	fmt.Scan(&user.Password)
	return user
}

func (c *consoleImpl) newUser() {
	user := CreateUser()

	out, err := c.register.SignUp(user)
	if err != nil {
		log.Println(err)
	} else {
		JSON(out)
	}
}

func (c *consoleImpl) confirmUser() {
	user := startConfirmation()

	out, err := c.register.ConfirmSignUp(user)
	if err != nil {
		log.Println(err)
	} else {
		JSON(out)
	}
}

func startConfirmation() *models.User {
	user := &models.User{}
	fmt.Print("Username? ")
	fmt.Scan(&user.Nickname)
	fmt.Print("Confimation Code? ")
	fmt.Scan(&user.ConfirmationCode)

	return user
}
