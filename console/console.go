package console

import (
	"custom-auth/cmd"
	"custom-auth/cognito"
	"custom-auth/email"
	"encoding/json"
	"fmt"
	"log"
)

type Console interface {
	Run()
}

type consoleImpl struct {
	register cognito.RegisterService
	admin    cognito.AdminService
	auth     cognito.LoginService
	emails   email.EmailService
}

const (
	CLEAR = "c"
)

func NewConsole(register cognito.RegisterService,
	admin cognito.AdminService,
	auth cognito.LoginService,
	emails email.EmailService) *consoleImpl {

	return &consoleImpl{
		register: register,
		admin:    admin,
		auth:     auth,
		emails:   emails,
	}
}

func JSON(data interface{}) {
	b, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(b))
	}
}

func (c *consoleImpl) Run() {
	c.menu()
}

func (c *consoleImpl) menu() {
	option := ""

	fmt.Println("==================================")
	fmt.Println("         AWS COGNITO CLI")
	fmt.Println("==================================")

	fmt.Println("[1] - Criar usuário")
	fmt.Println("[2] - Confirmar usuário")
	fmt.Println("[3] - Procurar usuário")
	fmt.Println("[4] - Logar")
	fmt.Println("[5] - Login personalizado")
	fmt.Println("[E] - Enviar email")
	fmt.Println("[V] - Verificar email")
	fmt.Println("[L] - Listar emails verificados")
	fmt.Println("[X] - Excluir usuário")
	fmt.Println("[C] - Limpar console")

	fmt.Println("")
	fmt.Print("Choice: ")
	fmt.Scan(&option)

	c.exec(option)
}

func (c *consoleImpl) exec(option string) {
	switch option {
	case CREATE_USER:
		c.newUser()
	case CONFIRM_USER:
		c.confirmUser()
	case FIND_USER:
		c.findUser()
	case LOGIN:
		c.authenticate()
	case CUSTOM_LOGIN:
		c.customAuth()
	case VERIFY_EMAIL:
		c.verifyEmail()
	case SEND_EMAIL:
		c.sendEmail()
	case LIST_EMAILS:
		c.listEmails()
	case DELETE_USER:
		c.deleteUser()
	case CLEAR:
		cmd.Clear()
	default:
		fmt.Println("Invalid option.")
	}
}
