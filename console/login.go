package console

import (
	"custom-auth/models"
	"fmt"
	"log"
	"strings"
)

const (
	LOGIN        = "4"
	CUSTOM_LOGIN = "5"
)

func (c *consoleImpl) authenticate() {
	user := &models.User{}

	fmt.Print("Username? ")
	fmt.Scan(&user.Nickname)
	fmt.Print("Password? ")
	fmt.Scan(&user.Password)

	c.auth.Login(user)
}

func (c *consoleImpl) customAuth() {
	user := &models.User{}

	fmt.Print("Username? ")
	fmt.Scan(&user.Nickname)

	out, err := c.auth.CustomLogin(user)
	if err != nil {
		log.Println(err)
		return
	}

	JSON(out)

	fmt.Print("Código de Confirmação: ")
	fmt.Scan(&user.ConfirmationCode)

	out2, err := c.auth.FinishCustomLogin(user, out)
	if err != nil {
		log.Println(err)
		return
	}

	JSON(out2)

	if *out2.ChallengeName == "" {
		fmt.Println("Logado com sucesso!")
		return
	}

	for {

		fmt.Println("Código Inválido!")
		fmt.Print("Insira o código de confirmação válido: ")
		fmt.Scan(&user.ConfirmationCode)

		out.Session = out2.Session

		out2, err = c.auth.FinishCustomLogin(user, out)

		if err != nil {
			log.Println(err)

			if strings.Contains(err.Error(), "NotAuthorizedException: Incorrect username or password.") {
				break
			}
			continue
		}

		JSON(out2)

		if out2.ChallengeName == nil {
			fmt.Println("Logado com sucesso!")
			break
		}
	}
}
