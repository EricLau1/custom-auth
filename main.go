package main

import (
	"custom-auth/cognito"
	"custom-auth/console"
	"custom-auth/email"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	provider := cognito.New()
	register := cognito.NewRegisterService(provider)
	admin := cognito.NewAdminService(provider)
	auth := cognito.NewLoginService(provider)
	emails := email.NewEmailService()

	cli := console.NewConsole(register, admin, auth, emails)

	run(cli)
}

func run(cli console.Console) {
	for {
		cli.Run()
		fmt.Println("")
	}
}
